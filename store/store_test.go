package store

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/Lagrange-Labs/lagrange-node/core"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
	"github.com/Lagrange-Labs/lagrange-node/store/mongodb"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
)

type StorageTestSuite struct {
	suite.Suite

	NewStorage func() (types.Storage, error)
}

func (s *StorageTestSuite) TestBatch() {
	storage, err := s.NewStorage()
	s.Require().NoError(err)
	chainID := uint32(1)

	batch1 := &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{
			BatchNumber:   1,
			L1BlockNumber: 1,
			L1TxHash:      core.RandomHex(32),
			ChainId:       chainID,
			L2Blocks: []*sequencerv2types.BlockHeader{
				{
					BlockNumber: 1,
					BlockHash:   core.RandomHex(32),
				},
				{
					BlockNumber: 2,
					BlockHash:   core.RandomHex(32),
				},
			},
		},
		CommitteeHeader: &sequencerv2types.CommitteeHeader{
			CurrentCommittee: core.RandomHex(32),
			NextCommittee:    core.RandomHex(32),
			TotalVotingPower: 100,
		},
		SequencedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
	}
	batch2 := &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{
			BatchNumber:   2,
			L1BlockNumber: 2,
			L1TxHash:      core.RandomHex(32),
			ChainId:       chainID,
			L2Blocks: []*sequencerv2types.BlockHeader{
				{
					BlockNumber: 3,
					BlockHash:   core.RandomHex(32),
				},
				{
					BlockNumber: 4,
					BlockHash:   core.RandomHex(32),
				},
			},
		},
		CommitteeHeader: &sequencerv2types.CommitteeHeader{
			CurrentCommittee: core.RandomHex(32),
			NextCommittee:    core.RandomHex(32),
			TotalVotingPower: 100,
		},
		SequencedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
	}

	ctx := context.Background()
	err = storage.AddBatch(ctx, batch1)
	s.Require().NoError(err)
	err = storage.AddBatch(ctx, batch2)
	s.Require().NoError(err)

	bn, err := storage.GetLastBatchNumber(ctx, chainID)
	s.Require().NoError(err)
	s.Require().Equal(batch2.BatchNumber(), bn)

	batch, err := storage.GetBatch(ctx, chainID, bn)
	s.Require().NoError(err)
	s.Require().Equal(batch.SequencedTime, batch2.SequencedTime)
	s.Require().Equal(batch.BatchHeader.L2Blocks[0], batch2.BatchHeader.L2Blocks[0])

	bnf, err := storage.GetLastFinalizedBatchNumber(ctx, chainID)
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), bnf)

	batch.AggSignature = core.RandomHex(96)
	batch.PubKeys = []string{core.RandomHex(32), core.RandomHex(32)}
	batch.FinalizedTime = time.Now().Format("2024-01-01 00:00:00.000000")

	err = storage.UpdateBatch(ctx, batch)
	s.Require().NoError(err)

	bn, err = storage.GetLastFinalizedBatchNumber(ctx, chainID)
	s.Require().NoError(err)
	s.Require().Equal(batch.BatchNumber(), bn)

	batch3, err := storage.GetBatch(ctx, chainID, bn)
	s.Require().NoError(err)
	s.Require().Equal(batch.FinalizedTime, batch3.FinalizedTime)
	s.Require().Equal(batch.AggSignature, batch3.AggSignature)
	s.Require().Equal(batch.PubKeys, batch3.PubKeys)

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
