package client

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/network"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
		ECDSAPrivateKey: "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429",
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

func (suite *ClientTestSuite) Test_Join_Network() {
	go suite.client.Start()
	time.Sleep(1 * time.Second)

	stakeAddress := suite.client.GetStakeAddress()
	node, err := suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress)
	suite.Require().NoError(err)
	suite.Require().Equal(networktypes.NodeJoined, node.Status)

	auth, err := utils.GetSigner(context.Background(), suite.ethClient, suite.cfg.ECDSAPrivateKey)
	suite.Require().NoError(err)
	err = testutil.RegisterOperator(suite.ethClient, auth, common.HexToAddress(stakeAddr), common.HexToAddress(slasherAddr))
	suite.Require().NoError(err)
	suite.manager.RunSequencer()
	time.Sleep(5 * time.Second)

	node, err = suite.manager.Storage.GetNodeByStakeAddr(context.Background(), stakeAddress)
	suite.Require().NoError(err)
	suite.Require().Equal(networktypes.NodeRegistered, node.Status)
	suite.client.Stop()
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
