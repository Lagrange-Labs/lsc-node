package integration

import (
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/store"
	"github.com/stretchr/testify/suite"
)

type CommitteeRootTestSuite struct {
	suite.Suite
	sequencer *sequencer.Sequencer
}

func (suite *CommitteeRootTestSuite) SetupTest() {
	var cfg, err = config.Default()
	if err != nil {
		panic(err)
	}
	store, err := store.NewStorage(&cfg.Store)
	if err != nil {
		panic(err)
	}

	suite.sequencer, err = sequencer.NewSequencer(&cfg.Sequencer, &cfg.RpcClient, store)
	if err != nil {
		panic(err)
	}
}

func (suite *CommitteeRootTestSuite) Test_Committee_root() {
	suite.sequencer.FetchCommitteeRoot(75)
	// cfg, err := config.Default()
	// require.NoError(suite.T(), err)

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		suite.T().Fatal("timeout")
	// 	default:
	// 	}
	// 	batch, err := suite.manager.Storage.GetBatch(ctx, suite.manager.GetChainID(), cfg.Sequencer.FromL1BlockNumber)
	// 	if errors.Is(err, types.ErrBatchNotFound) {
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// 	require.NoError(suite.T(), err)
	// 	require.NotNil(suite.T(), batch)
	// 	require.NotNil(suite.T(), batch.AggSignature)
	// 	break
	// }
}

func TestCommitteeRootTestSuite(t *testing.T) {
	suite.Run(t, new(CommitteeRootTestSuite))
}
