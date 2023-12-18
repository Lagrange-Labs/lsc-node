package integration

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil/operations"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			suite.T().Fatal("timeout")
		default:
		}
		block, err := suite.manager.Storage.GetBlock(ctx, suite.manager.GetChainID(), 85)
		if errors.Is(err, types.ErrBlockNotFound) {
			time.Sleep(1 * time.Second)
			continue
		}
		require.NoError(suite.T(), err)
		require.NotNil(suite.T(), block)
		require.NotNil(suite.T(), block.AggSignature)
		break
	}
}

func TestSequencerTestSuite(t *testing.T) {
	suite.Run(t, new(SequencerTestSuite))
}
