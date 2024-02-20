package integration

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/network"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
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
		BLSPrivateKey:   "0x00000000000000000000000000000000000000000000000000000000499602d4",
		ECDSAPrivateKey: "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429",
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
		require.NoError(t, suite.client.TryJoinNetwork())

		stakeAddress := suite.client.GetStakeAddress()
		node, err := suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress, suite.client.GetChainID())
		require.NoError(t, err)
		require.Equal(t, networktypes.NodeJoined, node.Status)

		suite.manager.RunSequencer(true)

		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Second)
			node, err = suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress, suite.client.GetChainID())
			if err == types.ErrNodeNotFound || node.Status == networktypes.NodeJoined {
				continue
			}
			require.NoError(t, err)
			require.Equal(t, networktypes.NodeRegistered, node.Status)
			break
		}
	})
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
