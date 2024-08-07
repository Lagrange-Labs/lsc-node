package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/Lagrange-Labs/lagrange-node/client"
	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

type ClientTestSuite struct {
	suite.Suite

	cfg     client.Config
	client  *client.Client
	manager *operations.Manager
}

func (suite *ClientTestSuite) SetupTest() {
	cfg, err := config.Default()
	suite.Require().NoError(err)

	dir := suite.T().TempDir()
	ecdsaKeyPath := dir + "/ecdsa.key"
	err = crypto.SaveKey("ECDSA", utils.Hex2Bytes("0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3"), "password", ecdsaKeyPath)
	suite.Require().NoError(err)
	blsKeyPath := dir + "/bls.key"
	err = crypto.SaveKey("BN254", utils.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000000499602d7"), "password", blsKeyPath)
	suite.Require().NoError(err)

	suite.cfg = client.Config{
		GrpcURLs:    []string{"127.0.0.1:9090"},
		Chain:       "mock",
		EthereumURL: "http://localhost:8545",
		// BLSKeystorePath:             blsKeyPath,
		// BLSKeystorePassword:         "password",
		// SignerECDSAKeystorePath:     ecdsaKeyPath,
		// SignerECDSAKeystorePassword: "password",
		OperatorAddress:    "0x13cF11F76a08214A826355a1C8d661E41EA7Bf97",
		CommitteeSCAddress: cfg.Client.CommitteeSCAddress,
		PullInterval:       utils.TimeDuration(2 * time.Second),
		BLSCurve:           "BN254",
	}
	suite.manager, err = operations.NewManager()
	suite.Require().NoError(err)
	suite.manager.RunServer()
	time.Sleep(1 * time.Second)
	suite.client, err = client.NewClient(&suite.cfg, suite.manager.GetRpcConfig())
	suite.Require().NoError(err)
}

func (suite *ClientTestSuite) TearDownSuite() {
	suite.client.Stop()
	suite.manager.Close()
}

func (suite *ClientTestSuite) Test_Client_Start() {
	suite.T().Run("Test_Join_Network", func(t *testing.T) {
		suite.manager.RunSequencer(true)
		err := suite.client.TryJoinNetwork()
		require.NoError(t, err)

		stakeAddress := suite.client.GetStakeAddress()
		node, err := suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress, suite.manager.GetChainID())
		require.NoError(t, err)
		require.Equal(t, servertypes.NodeJoined, node.Status)
	})
}

func TestClientTestSuite(t *testing.T) {
	// suite.Run(t, new(ClientTestSuite))
}
