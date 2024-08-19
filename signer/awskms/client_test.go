package awskms

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/signer"
)

const localStackPort = "4566"

type TestClientSuite struct {
	suite.Suite

	client    *Client
	container testcontainers.Container
}

func (s *TestClientSuite) SetupSuite() {
	container, err := startLocalstackContainer("get_public_key_test")
	require.NoError(s.T(), err)
	s.container = container
	mappedPort, err := container.MappedPort(context.Background(), localStackPort)
	require.NoError(s.T(), err)

	cfg := &signer.AWSKMSConfig{
		Region:          "us-east-1",
		EndpointURL:     fmt.Sprintf("http://127.0.0.1:%s", mappedPort),
		AccessKeyID:     "localstack",
		SecretAccessKey: "localstack",
		KeyID:           "localstack",
	}
	s.client, err = NewClient(cfg)
	require.NoError(s.T(), err)

	res, err := s.client.CreateKey(context.Background(), &kms.CreateKeyInput{
		KeySpec:  types.KeySpecEccSecgP256k1,
		KeyUsage: types.KeyUsageTypeSignVerify,
	})
	require.NoError(s.T(), err)
	s.client.keyID = *res.KeyMetadata.KeyId

}

func (s *TestClientSuite) TearDownSuite() {
	require.NoError(s.T(), s.container.Terminate(context.Background()))
}

func startLocalstackContainer(name string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image: "localstack/localstack:latest",
		Name:  fmt.Sprintf("localstack-test-%s", name),
		Env: map[string]string{
			"LOCALSTACK_HOST": fmt.Sprintf("localhost.localstack.cloud:%s", localStackPort),
		},
		ExposedPorts: []string{localStackPort},
		WaitingFor:   wait.ForLog("Ready."),
		AutoRemove:   true,
	}
	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func (s *TestClientSuite) TestSign() {
	msg := crypto.Hash([]byte("hello"))
	sig, err := s.client.Sign(msg, false)
	require.NoError(s.T(), err)

	verified, _, err := crypto.VerifyECDSASignature(msg, sig)
	require.NoError(s.T(), err)
	require.True(s.T(), verified)
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}
