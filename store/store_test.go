package store

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
	"github.com/Lagrange-Labs/lagrange-node/store/mongodb"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

type StorageTestSuite struct {
	suite.Suite

	NewStorage func() (types.Storage, error)
}

func (s *StorageTestSuite) TestBatch() {
	storage, err := s.NewStorage()
	s.Require().NoError(err)
	chainID := uint32(1)

	batch := &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{
			BatchNumber:   1,
			BatchHash:     utils.RandomHex(32),
			L1BlockNumber: 1,
			L1TxHash:      utils.RandomHex(32),
			ChainId:       chainID,
			L2Blocks: []*sequencerv2types.BlockHeader{
				{
					BlockNumber: 1,
					BlockHash:   utils.RandomHex(32),
				},
				{
					BlockNumber: 2,
					BlockHash:   utils.RandomHex(32),
				},
			},
		},
		CommitteeHeader: &sequencerv2types.CommitteeHeader{
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
			TotalVotingPower: 100,
		},
		SequencedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
	}

	ctx := context.Background()
	err = storage.AddBatch(ctx, batch)
	s.Require().NoError(err)

	bn, err := storage.GetLastBatchNumber(ctx, chainID)
	s.Require().NoError(err)
	s.Require().Equal(batch.BatchNumber(), bn)

	batch2, err := storage.GetBatch(ctx, chainID, bn)
	s.Require().NoError(err)
	s.Require().Equal(batch.SequencedTime, batch2.SequencedTime)
	s.Require().Equal(batch.BatchHeader.L2Blocks[0], batch2.BatchHeader.L2Blocks[0])

	batch.AggSignature = utils.RandomHex(96)
	batch.PubKeys = []string{utils.RandomHex(32), utils.RandomHex(32)}
	batch.FinalizedTime = time.Now().Format("2024-01-01 00:00:00.000000")

	err = storage.UpdateBatch(ctx, batch)
	s.Require().NoError(err)

	bn, err = storage.GetLastFinalizedBatchNumber(ctx, chainID)
	s.Require().NoError(err)
	s.Require().Equal(batch.BatchNumber(), bn)

	batch2, err = storage.GetBatch(ctx, chainID, bn)
	s.Require().NoError(err)
	s.Require().Equal(batch.FinalizedTime, batch2.FinalizedTime)
	s.Require().Equal(batch.AggSignature, batch2.AggSignature)
	s.Require().Equal(batch.PubKeys, batch2.PubKeys)

	// clean up
	err = storage.CleanUp(ctx)
	s.Require().NoError(err)
}

func TestMemDBSuit(t *testing.T) {
	suite.Run(t, &StorageTestSuite{
		NewStorage: func() (types.Storage, error) {
			return memdb.NewMemDB()
		},
	})
}

func TestMongoDBSuit(t *testing.T) {
	suite.Run(t, &StorageTestSuite{
		NewStorage: func() (types.Storage, error) {
			return mongodb.NewMongoDB("mongodb://127.0.0.1:27017")
		},
	})
}
