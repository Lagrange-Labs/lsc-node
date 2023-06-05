package client

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/network"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	stakeAddr   = "0x98f07aB2d35638B79582b250C01444cEce0E517A"
	slasherAddr = "0x6Bf0fF4eBa00E3668c0241bb1C622CDBFE55bbE0"
)

type ClientTestSuite struct {
	suite.Suite

	cfg       network.ClientConfig
	client    *network.Client
	ethClient *ethclient.Client
	manager   *operations.Manager
}

func (suite *ClientTestSuite) SetupTest() {
	suite.cfg = network.ClientConfig{
		GrpcURL:         "127.0.0.1:9090",
		Chain:           "arbitrum",
		RPCEndpoint:     "http://127.0.0.1:8545",
		BLSPrivateKey:   "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a",
		ECDSAPrivateKey: "0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3",
		PullInterval:    utils.TimeDuration(2 * time.Second),
	}
	var err error
	suite.manager, err = operations.NewManager()
	suite.Require().NoError(err)
	suite.manager.RunServer()
	time.Sleep(1 * time.Second)
	suite.client, err = network.NewClient(&suite.cfg)
	suite.Require().NoError(err)
	suite.ethClient, err = ethclient.Dial(suite.cfg.RPCEndpoint)
	suite.Require().NoError(err)
}

func (suite *ClientTestSuite) TearDownSuite() {
	suite.client.Stop()
}

func (suite *ClientTestSuite) Test_Client_Start() {
	var (
		block *sequencertypes.Block
		err   error
	)

	suite.T().Run("Test_Join_Network", func(t *testing.T) {
		require.NoError(t, suite.client.TryJoinNetwork())

		stakeAddress := suite.client.GetStakeAddress()
		node, err := suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress)
		require.NoError(t, err)
		require.Equal(t, networktypes.NodeJoined, node.Status)

		auth, err := utils.GetSigner(context.Background(), suite.ethClient, suite.cfg.ECDSAPrivateKey)
		require.NoError(t, err)
		err = testutil.RegisterOperator(suite.ethClient, auth, common.HexToAddress(stakeAddr), common.HexToAddress(slasherAddr))
		require.NoError(t, err)
		suite.manager.RunSequencer()
		time.Sleep(3 * time.Second)

		node, err = suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress)
		require.NoError(t, err)
		require.Equal(t, networktypes.NodeRegistered, node.Status)
	})

	suite.T().Run("Test_Get_Block", func(t *testing.T) {
		block, err = suite.client.TryGetBlock()
		if err == network.ErrBlockNotReady {
			time.Sleep(3 * time.Second)
			block, err = suite.client.TryGetBlock()
		}
		require.NoError(t, err)
		require.NotNil(t, block)
	})

	suite.T().Run("Test_Commit_Block", func(t *testing.T) {
		err = suite.client.TryCommitBlock(block)
		require.NoError(t, err)
	})
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
