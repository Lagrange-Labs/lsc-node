package mongodb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

func TestConvertProtobufToMongo(t *testing.T) {
	block := &sequencertypes.Block{
		BlockHeader: &sequencertypes.BlockHeader{
			CurrentCommittee:  utils.RandomHex(32),
			NextCommittee:     utils.RandomHex(32),
			ProposerPubKey:    utils.RandomHex(32),
			ProposerSignature: utils.RandomHex(32),
		},
		ChainHeader: &sequencertypes.ChainHeader{
			ChainId:       1,
			BlockHash:     utils.RandomHex(32),
			BlockNumber:   1,
			L1BlockNumber: 1,
			L1TxHash:      utils.RandomHex(32),
		},
		PubKeys: []string{
			utils.RandomHex(32),
			utils.RandomHex(32),
		},
		AggSignature:  utils.RandomHex(96),
		SequencedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
		FinalizedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
	}
	mBlock, err := ConvertProtobufToMongo(block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader.CurrentCommittee, mBlock["block_header"].(primitive.M)["current_committee"])
	require.Equal(t, block.BlockHeader.ProposerPubKey, mBlock["block_header"].(primitive.M)["proposer_pub_key"])
	require.Equal(t, block.ChainHeader.BlockHash, mBlock["chain_header"].(primitive.M)["block_hash"])
	require.Equal(t, block.L1BlockNumber(), mBlock["chain_header"].(primitive.M)["l1_block_number"])
	require.Equal(t, block.BlockNumber(), mBlock["chain_header"].(primitive.M)["block_number"])
	require.Equal(t, len(block.PubKeys), len(mBlock["pub_keys"].(primitive.A)))
	require.Equal(t, block.PubKeys[0], mBlock["pub_keys"].(primitive.A)[0].(string))
	require.Equal(t, block.AggSignature, mBlock["agg_signature"])
	require.Equal(t, block.SequencedTime, mBlock["sequenced_time"])
	require.Equal(t, block.FinalizedTime, mBlock["finalized_time"])

	// Test with batch
	batch := &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{
			BatchNumber:   1,
			L1BlockNumber: 1,
			L1TxHash:      utils.RandomHex(32),
			L1TxIndex:     1,
			ChainId:       1,
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
		PubKeys: []string{
			utils.RandomHex(32),
			utils.RandomHex(32),
		},
		AggSignature:  utils.RandomHex(96),
		SequencedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
		FinalizedTime: time.Now().Format("2024-01-01 00:00:00.000000"),
	}
	mBatch, err := ConvertProtobufToMongo(batch)
	require.NoError(t, err)
	require.Equal(t, batch.BatchHeader.BatchNumber, mBatch["batch_header"].(primitive.M)["batch_number"])
	require.Equal(t, batch.BatchHeader.L1BlockNumber, mBatch["batch_header"].(primitive.M)["l1_block_number"])
	require.Equal(t, batch.BatchHeader.L1TxHash, mBatch["batch_header"].(primitive.M)["l1_tx_hash"])
	require.Equal(t, batch.BatchHeader.L1TxIndex, mBatch["batch_header"].(primitive.M)["l1_tx_index"])
	require.Equal(t, batch.BatchHeader.ChainId, mBatch["batch_header"].(primitive.M)["chain_id"])
	require.Equal(t, len(batch.BatchHeader.L2Blocks), len(mBatch["batch_header"].(primitive.M)["l2_blocks"].(primitive.A)))
	require.Equal(t, batch.BatchHeader.L2Blocks[0].BlockNumber, mBatch["batch_header"].(primitive.M)["l2_blocks"].(primitive.A)[0].(primitive.M)["block_number"])
	require.Equal(t, batch.BatchHeader.L2Blocks[1].BlockHash, mBatch["batch_header"].(primitive.M)["l2_blocks"].(primitive.A)[1].(primitive.M)["block_hash"])
	require.Equal(t, batch.CommitteeHeader.CurrentCommittee, mBatch["committee_header"].(primitive.M)["current_committee"])
	require.Equal(t, batch.CommitteeHeader.NextCommittee, mBatch["committee_header"].(primitive.M)["next_committee"])
	require.Equal(t, batch.CommitteeHeader.TotalVotingPower, mBatch["committee_header"].(primitive.M)["total_voting_power"])
	require.Equal(t, len(batch.PubKeys), len(mBatch["pub_keys"].(primitive.A)))
	require.Equal(t, batch.PubKeys[0], mBatch["pub_keys"].(primitive.A)[0].(string))
	require.Equal(t, batch.AggSignature, mBatch["agg_signature"])
	require.Equal(t, batch.SequencedTime, mBatch["sequenced_time"])
	require.Equal(t, batch.FinalizedTime, mBatch["finalized_time"])
}

func TestConvertMongoToBlock(t *testing.T) {
	m := bson.M{
		"block_header": bson.M{
			"current_committee":  utils.RandomHex(32),
			"next_committee":     utils.RandomHex(32),
			"proposer_pub_key":   utils.RandomHex(32),
			"proposer_signature": utils.RandomHex(32),
		},
		"chain_header": bson.M{
			"chain_id":        int64(1),
			"block_hash":      utils.RandomHex(32),
			"block_number":    int64(1),
			"l1_block_number": int64(1),
			"l1_tx_hash":      utils.RandomHex(32),
		},
		"pub_keys": bson.A{
			utils.RandomHex(32),
			utils.RandomHex(32),
		},
		"agg_signature":  utils.RandomHex(96),
		"sequenced_time": time.Now().Format("2024-01-01 00:00:00.000000"),
		"finalized_time": time.Now().Format("2024-01-01 00:00:00.000000"),
	}
	block := ConvertMongoToBlock(m)
	require.Equal(t, m["block_header"].(bson.M)["current_committee"], block.BlockHeader.CurrentCommittee)
	require.Equal(t, m["block_header"].(bson.M)["proposer_pub_key"], block.BlockHeader.ProposerPubKey)
	require.Equal(t, m["chain_header"].(bson.M)["block_hash"], block.ChainHeader.BlockHash)
	require.Equal(t, m["chain_header"].(bson.M)["block_number"], int64(block.BlockNumber()))
	require.Equal(t, m["chain_header"].(bson.M)["l1_block_number"], int64(block.L1BlockNumber()))
	require.Equal(t, len(m["pub_keys"].(bson.A)), len(block.PubKeys))
	require.Equal(t, m["pub_keys"].(bson.A)[0], block.PubKeys[0])
	require.Equal(t, m["pub_keys"].(bson.A)[1], block.PubKeys[1])
	require.Equal(t, m["agg_signature"], block.AggSignature)
	require.Equal(t, m["sequenced_time"], block.SequencedTime)
	require.Equal(t, m["finalized_time"], block.FinalizedTime)
}

func TestConvertMongoToBatch(t *testing.T) {
	m := bson.M{
		"batch_header": bson.M{
			"batch_number":    int64(1),
			"l1_block_number": int64(1),
			"l1_tx_hash":      utils.RandomHex(32),
			"l1_tx_index":     int64(1),
			"chain_id":        int64(1),
			"l2_blocks": bson.A{
				bson.M{
					"block_number": int64(1),
					"block_hash":   utils.RandomHex(32),
				},
				bson.M{
					"block_number": int64(2),
					"block_hash":   utils.RandomHex(32),
				},
			},
		},
		"committee_header": bson.M{
			"current_committee":  utils.RandomHex(32),
			"next_committee":     utils.RandomHex(32),
			"total_voting_power": int64(100),
		},
		"pub_keys":       bson.A{utils.RandomHex(32), utils.RandomHex(32)},
		"agg_signature":  utils.RandomHex(96),
		"sequenced_time": time.Now().Format("2024-01-01 00:00:00.000000"),
		"finalized_time": time.Now().Format("2024-01-01 00:00:00.000000"),
	}
	batch := ConvertMongoToBatch(m)
	require.Equal(t, m["batch_header"].(bson.M)["batch_number"], int64(batch.BatchNumber()))
	require.Equal(t, m["batch_header"].(bson.M)["l1_block_number"], int64(batch.L1BlockNumber()))
	require.Equal(t, m["batch_header"].(bson.M)["l1_tx_hash"], batch.L1TxHash())
	require.Equal(t, m["batch_header"].(bson.M)["l1_tx_index"], int64(batch.BatchHeader.L1TxIndex))
	require.Equal(t, m["batch_header"].(bson.M)["chain_id"], int64(batch.ChainID()))
	require.Equal(t, len(m["batch_header"].(bson.M)["l2_blocks"].(bson.A)), len(batch.BatchHeader.L2Blocks))
	require.Equal(t, m["batch_header"].(bson.M)["l2_blocks"].(bson.A)[0].(bson.M)["block_number"], int64(batch.BatchHeader.L2Blocks[0].BlockNumber))
	require.Equal(t, m["batch_header"].(bson.M)["l2_blocks"].(bson.A)[1].(bson.M)["block_hash"], batch.BatchHeader.L2Blocks[1].BlockHash)
	require.Equal(t, m["committee_header"].(bson.M)["current_committee"], batch.CommitteeHeader.CurrentCommittee)
	require.Equal(t, m["committee_header"].(bson.M)["next_committee"], batch.CommitteeHeader.NextCommittee)
	require.Equal(t, m["committee_header"].(bson.M)["total_voting_power"], int64(batch.CommitteeHeader.TotalVotingPower))
	require.Equal(t, len(m["pub_keys"].(bson.A)), len(batch.PubKeys))
	require.Equal(t, m["pub_keys"].(bson.A)[0], batch.PubKeys[0])
	require.Equal(t, m["pub_keys"].(bson.A)[1], batch.PubKeys[1])
	require.Equal(t, m["agg_signature"], batch.AggSignature)
	require.Equal(t, m["sequenced_time"], batch.SequencedTime)
	require.Equal(t, m["finalized_time"], batch.FinalizedTime)
}
