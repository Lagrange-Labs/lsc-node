package integration

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SequencerTestSuite struct {
	suite.Suite

	cfg       sequencer.Config
	sequencer *sequencer.Sequencer
	ethClient *ethclient.Client
	manager   *operations.Manager
}

func (suite *SequencerTestSuite) SetupTest() {
	suite.cfg = sequencer.Config{
		Chain:           "arbitrum",
		RPCURL:          "http://localhost:8545",
		EthURL:          "http://localhost:8545",
		FromBlockNumber: 1,
	}

	var err error
	suite.manager, err = operations.NewManager()
	suite.Require().NoError(err)
	suite.manager.RunServer()
	time.Sleep(1 * time.Second)
	suite.manager.RunClients()
	suite.sequencer, err = sequencer.NewSequencer(&suite.cfg, suite.manager.Storage)
	suite.Require().NoError(err)
	suite.ethClient, err = ethclient.Dial(suite.cfg.RPCURL)
	suite.Require().NoError(err)
}

func (suite *SequencerTestSuite) TearDownSuite() {
	suite.sequencer.Stop()
}

func (suite *SequencerTestSuite) Test_Sequencer_Block_Generation() {
	go suite.sequencer.Start() // nolint:errcheck
	time.Sleep(5 * time.Second)

	block, err := suite.manager.Storage.GetBlock(context.Background(), suite.sequencer.GetChainID(), 150)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), block)
	require.NotNil(suite.T(), block.AggSignature)

	suite.sequencer.Stop()
}

func TestSequencerTestSuite(t *testing.T) {
	suite.Run(t, new(SequencerTestSuite))
}
