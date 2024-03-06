package integration

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/network"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite

	cfg     network.ClientConfig
	client  *network.Client
	manager *operations.Manager
}

func (suite *ClientTestSuite) SetupTest() {
	suite.cfg = network.ClientConfig{
		GrpcURL:         "127.0.0.1:9090",
		Chain:           "arbitrum",
		EthereumURL:     "http://localhost:8545",
		BLSPrivateKey:   "0x00000000000000000000000000000000000000000000000000000000499602d7",
		ECDSAPrivateKey: "0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3",
		PullInterval:    utils.TimeDuration(2 * time.Second),
		BLSCurve:        "BN254",
	}
	var err error
	suite.manager, err = operations.NewManager()
	suite.Require().NoError(err)
	suite.manager.RunServer()
	time.Sleep(1 * time.Second)
	suite.client, err = network.NewClient(&suite.cfg, suite.manager.GetRpcConfig())
	suite.Require().NoError(err)
}

func (suite *ClientTestSuite) TearDownSuite() {
	suite.client.Stop()
	suite.manager.Close()
}

func (suite *ClientTestSuite) Test_Client_Start() {
	suite.T().Run("Test_Join_Network", func(t *testing.T) {
		suite.manager.RunSequencer(true)
		suite.client.TryJoinNetwork()

		stakeAddress := suite.client.GetStakeAddress()
		node, err := suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress, suite.client.GetChainID())
		require.NoError(t, err)
		require.Equal(t, networktypes.NodeJoined, node.Status)
	})
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
