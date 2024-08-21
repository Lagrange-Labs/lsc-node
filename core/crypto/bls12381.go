package crypto

import (
	"fmt"

	blst "github.com/supranational/blst/bindings/go"
	"github.com/umbracle/go-eth-consensus/bls"
)

// Scheme is the crypto scheme implementation for BLS12-381 curve.
type BLS12381Scheme struct {
}

var _ BLSScheme = (*BLS12381Scheme)(nil)

func (s *BLS12381Scheme) VerifySignature(pubKey, message, signature []byte, _ bool) (bool, error) {
	pubK := new(bls.PublicKey)
	if err := pubK.Deserialize(pubKey); err != nil {
		return false, err
	}
	sig := new(bls.Signature)
	if err := sig.Deserialize(signature); err != nil {
		return false, err
	}
	return sig.VerifyByte(pubK, message)
}

func (s *BLS12381Scheme) AggregateSignatures(signatures [][]byte, _ bool) ([]byte, error) {
	sigs := make([]*bls.Signature, len(signatures))
	for i, sig := range signatures {
		s := new(bls.Signature)
		if err := s.Deserialize(sig); err != nil {
			return nil, err
		}
		sigs[i] = s
	}
	aggSig := bls.AggregateSignatures(sigs).Serialize()
	return aggSig[:], nil
}

func (s *BLS12381Scheme) aggregatePublicKeys(pubKeys [][]byte) (*blst.P1Affine, error) {
	pks := make([]*blst.P1Affine, len(pubKeys))
	for i, pk := range pubKeys {
		pub := new(blst.P1Affine).Uncompress(pk)
		if pub == nil {
			return nil, fmt.Errorf("failed to deserialize the public key")
		}
		if !pub.KeyValidate() {
			return nil, fmt.Errorf("public key not in group")
		}
		pks[i] = pub
	}
	aggregator := new(blst.P1Aggregate)
	if !aggregator.Aggregate(pks, false) {
		return nil, fmt.Errorf("failed to aggregate public keys")
	}

	return aggregator.ToAffine(), nil
}

func (s *BLS12381Scheme) AggregatePublicKeys(pubKeys [][]byte, _ bool) ([]byte, error) {
	aggPk, err := s.aggregatePublicKeys(pubKeys)
	if err != nil {
		return nil, err
	}

	aggPkRaw := aggPk.Compress()
	return aggPkRaw[:], nil
}

func (s *BLS12381Scheme) VerifyAggregatedSignature(pubKeys [][]byte, message, signature []byte, _ bool) (bool, error) {
	sig := new(bls.Signature)
	if err := sig.Deserialize(signature); err != nil {
		return false, err
	}

	pks := make([]*bls.PublicKey, len(pubKeys))
	for i, pk := range pubKeys {
		pub := new(bls.PublicKey)
		if err := pub.Deserialize(pk); err != nil {
			return false, err
		}
		pks[i] = pub
	}

	return sig.FastAggregateVerify(pks, message)
}

func (s *BLS12381Scheme) Sign(privKey, message []byte, _ bool) ([]byte, error) {
	priv := new(bls.SecretKey)
	if err := priv.Unmarshal(privKey); err != nil {
		return nil, err
	}
	sig, err := priv.Sign(message)
	if err != nil {
		return nil, err
	}
	sigRaw := sig.Serialize()
	return sigRaw[:], nil
}

func (s *BLS12381Scheme) GenerateRandomKey() ([]byte, error) {
	priv := bls.RandomKey()
	privRaw, err := priv.Marshal()
	if err != nil {
		return nil, err
	}
	return privRaw[:], nil
}

func (s *BLS12381Scheme) GetPublicKey(privKey []byte, isCompressed bool, _ bool) ([]byte, error) {
	priv := new(bls.SecretKey)
	if err := priv.Unmarshal(privKey); err != nil {
		return nil, err
	}

	pubRaw := priv.GetPublicKey().Serialize()
	if isCompressed {
		return pubRaw[:], nil
	}

	pubKey := new(blst.P1Affine).Uncompress(pubRaw[:])
	return pubKey.Serialize(), nil
}

func (s *BLS12381Scheme) ConvertPublicKey(pubKey []byte, isCompressed bool, _ bool) ([]byte, error) {
	if isCompressed {
		pubKey = new(blst.P1Affine).Deserialize(pubKey).Compress()

	} else {
		pubKey = new(blst.P1Affine).Uncompress(pubKey).Serialize()
	}

	return pubKey, nil
}
