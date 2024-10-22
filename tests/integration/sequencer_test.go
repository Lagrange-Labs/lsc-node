package integration

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lsc-node/config"
	"github.com/Lagrange-Labs/lsc-node/store/types"
	"github.com/Lagrange-Labs/lsc-node/testutil/operations"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SequencerTestSuite struct {
	suite.Suite

	manager *operations.Manager
}

func (suite *SequencerTestSuite) SetupTest() {
	var err error
	suite.manager, err = operations.NewManager()
	suite.Require().NoError(err)
	suite.manager.RunSequencer(false)
}

func (suite *SequencerTestSuite) TearDownSuite() {
	suite.manager.Close()
}

func (suite *SequencerTestSuite) Test_Sequencer_Block_Generation() {
	cfg, err := config.Default()
	require.NoError(suite.T(), err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			suite.T().Fatal("timeout")
		default:
		}
		batch, err := suite.manager.Storage.GetBatch(ctx, suite.manager.GetChainID(), cfg.Sequencer.FromL1BlockNumber)
		if errors.Is(err, types.ErrBatchNotFound) {
			time.Sleep(1 * time.Second)
			continue
		}
		require.NoError(suite.T(), err)
		require.NotNil(suite.T(), batch)
		require.NotNil(suite.T(), batch.AggSignature)
		break
	}
}

func TestSequencerTestSuite(t *testing.T) {
	suite.Run(t, new(SequencerTestSuite))
}
