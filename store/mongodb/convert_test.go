package mongodb

import (
	"fmt"
	"testing"
	"time"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
		FinalizedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
	}
	mBlock, err := ConvertProtobufToMongo(block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader.CurrentCommittee, mBlock["block_header"].(primitive.M)["current_committee"])
	require.Equal(t, block.BlockHeader.ProposerPubKey, mBlock["block_header"].(primitive.M)["proposer_pub_key"])
	require.Equal(t, block.ChainHeader.BlockHash, mBlock["chain_header"].(primitive.M)["block_hash"])
	require.Equal(t, block.L1BlockNumber(), mBlock["chain_header"].(primitive.M)["l1_block_number"])
	require.Equal(t, block.BlockNumber(), mBlock["chain_header"].(primitive.M)["block_number"])
	require.Equal(t, len(block.PubKeys), len(mBlock["pub_keys"].([]string)))
	require.Equal(t, block.PubKeys[0], mBlock["pub_keys"].([]string)[0])
	require.Equal(t, block.PubKeys[1], mBlock["pub_keys"].([]string)[1])
	require.Equal(t, block.AggSignature, mBlock["agg_signature"])
	require.Equal(t, block.SequencedTime, mBlock["sequenced_time"])
	require.Equal(t, block.FinalizedTime, mBlock["finalized_time"])
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
		"sequenced_time": fmt.Sprintf("%d", time.Now().UnixMicro()),
		"finalized_time": fmt.Sprintf("%d", time.Now().UnixMicro()),
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
