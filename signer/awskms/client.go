package awskms

import (
	"bytes"
	"context"
	"encoding/asn1"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/signer"
)

var secp256k1N = crypto.S256().Params().N
var secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))

var _ signer.Signer = (*Client)(nil)

// Client is a AWS KMS client.
type Client struct {
	*kms.Client

	keyID string
}

// NewClient creates a new AWS KMS client.
// It only supports ECDSA P-256.
func NewClient(cfg *signer.AWSKMSConfig) (*Client, error) {
	createClient := func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if cfg.EndpointURL != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           cfg.EndpointURL,
				SigningRegion: region,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}
	customResolver := aws.EndpointResolverWithOptionsFunc(createClient)

	awsCfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRetryMode(aws.RetryModeStandard),
	)
	if err != nil {
		return nil, err
	}
	if len(cfg.EndpointURL) == 0 {
		awsCfg.BaseEndpoint = aws.String(cfg.EndpointURL)
	}

	return &Client{
		Client: kms.NewFromConfig(awsCfg),
		keyID:  cfg.KeyID,
	}, nil
}

type asn1EcSig struct {
	R asn1.RawValue
	S asn1.RawValue
}

// Sign signs the data.
func (c *Client) Sign(digest []byte) ([]byte, error) {
	input := &kms.SignInput{
		KeyId:            &c.keyID,
		Message:          digest,
		MessageType:      types.MessageTypeDigest,
		SigningAlgorithm: types.SigningAlgorithmSpecEcdsaSha256,
	}

	output, err := c.Client.Sign(context.Background(), input)
	if err != nil {
		return nil, err
	}

	var sigAsn1 asn1EcSig
	_, err = asn1.Unmarshal(output.Signature, &sigAsn1)
	if err != nil {
		return nil, err
	}
	rBytes := sigAsn1.R.Bytes
	sBytes := sigAsn1.S.Bytes
	// Adjust S value from signature according to Ethereum standard
	sBigInt := new(big.Int).SetBytes(sBytes)
	if sBigInt.Cmp(secp256k1HalfN) > 0 {
		sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
	}
	rsSignature := append(adjustSignatureLength(rBytes), adjustSignatureLength(sBytes)...)

	return append(rsSignature, []byte{0}...), nil
}

type asn1EcPublicKey struct {
	EcPublicKeyInfo asn1EcPublicKeyInfo
	PublicKey       asn1.BitString
}

type asn1EcPublicKeyInfo struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.ObjectIdentifier
}

// GetPubKey gets the public key.
func (c *Client) GetPubKey() ([]byte, error) {
	output, err := c.Client.GetPublicKey(context.Background(), &kms.GetPublicKeyInput{KeyId: aws.String(c.keyID)})
	if err != nil {
		return nil, err
	}

	var pubKeyAsn1 asn1EcPublicKey
	_, err = asn1.Unmarshal(output.PublicKey, &pubKeyAsn1)
	if err != nil {
		return nil, err
	}

	pubkey, err := crypto.UnmarshalPubkey(pubKeyAsn1.PublicKey.Bytes)
	if err != nil {
		return nil, err
	}

	keyAddr := crypto.PubkeyToAddress(*pubkey)

	return keyAddr.Bytes(), nil
}

func adjustSignatureLength(buffer []byte) []byte {
	buffer = bytes.TrimLeft(buffer, "\x00")
	for len(buffer) < 32 {
		zeroBuf := []byte{0}
		buffer = append(zeroBuf, buffer...)
	}
	return buffer
}
