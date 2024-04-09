package mongodb

import (
	"fmt"
	"reflect"
	"strings"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertProtobufToMongo converts a protobuf object to a mongo object.
func ConvertProtobufToMongo(obj interface{}) (bson.M, error) {
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", value.Kind())
	}

	m := bson.M{}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			continue
		}
		if !fieldValue.IsValid() {
			continue
		}

		fieldName = strings.Split(fieldName, ",")[0]

		if fieldValue.Kind() == reflect.Struct {
			inner, err := ConvertProtobufToMongo(fieldValue.Interface())
			if err != nil {
				return nil, err
			}

			m[fieldName] = inner
		} else if fieldValue.Kind() == reflect.Slice {
			a := bson.A{}
			for j := 0; j < fieldValue.Len(); j++ {
				ele := fieldValue.Index(j)
				if ele.Kind() == reflect.Ptr {
					ele = ele.Elem()
				}
				if ele.Kind() == reflect.Struct {
					inner, err := ConvertProtobufToMongo(fieldValue.Index(j).Interface())
					if err != nil {
						return nil, err
					}

					a = append(a, inner)
				} else {
					a = append(a, fieldValue.Index(j).Interface())
				}
			}

			m[fieldName] = a
		} else {
			m[fieldName] = fieldValue.Interface()
		}
	}

	return m, nil
}

// ConvertMongoToBlock converts a mongo object to a protobuf block object.
func ConvertMongoToBlock(m bson.M) *sequencertypes.Block {
	block := &sequencertypes.Block{}
	if chainHeader, ok := m["chain_header"]; ok {
		mChainHeader := chainHeader.(bson.M)
		block.ChainHeader = &sequencertypes.ChainHeader{}
		block.ChainHeader.BlockHash = mChainHeader["block_hash"].(string)
		block.ChainHeader.BlockNumber = uint64(mChainHeader["block_number"].(int64))
		block.ChainHeader.L1BlockNumber = uint64(mChainHeader["l1_block_number"].(int64))
		block.ChainHeader.L1TxHash = mChainHeader["l1_tx_hash"].(string)
		block.ChainHeader.ChainId = uint32(mChainHeader["chain_id"].(int64))
	}

	if blockHeader, ok := m["block_header"]; ok {
		mBlockHeader := blockHeader.(bson.M)
		block.BlockHeader = &sequencertypes.BlockHeader{}
		block.BlockHeader.CurrentCommittee = mBlockHeader["current_committee"].(string)
		block.BlockHeader.NextCommittee = mBlockHeader["next_committee"].(string)
		block.BlockHeader.ProposerPubKey = mBlockHeader["proposer_pub_key"].(string)
		block.BlockHeader.ProposerSignature = mBlockHeader["proposer_signature"].(string)
	}

	block.AggSignature = m["agg_signature"].(string)
	if len(block.AggSignature) > 0 {
		block.PubKeys = make([]string, 0)
		for _, pubKey := range m["pub_keys"].(primitive.A) {
			block.PubKeys = append(block.PubKeys, pubKey.(string))
		}
	}

	block.SequencedTime = m["sequenced_time"].(string)
	block.FinalizedTime = m["finalized_time"].(string)

	return block
}

// ConvertMongoToBatch converts a mongo object to a protobuf batch object.
func ConvertMongoToBatch(m bson.M) *sequencerv2types.Batch {
	batch := &sequencerv2types.Batch{}
	if batchHeader, ok := m["batch_header"]; ok {
		mBatchHeader := batchHeader.(bson.M)
		batch.BatchHeader = &sequencerv2types.BatchHeader{}
		batch.BatchHeader.BatchNumber = uint64(mBatchHeader["batch_number"].(int64))
		batch.BatchHeader.L1BlockNumber = uint64(mBatchHeader["l1_block_number"].(int64))
		batch.BatchHeader.L1TxHash = mBatchHeader["l1_tx_hash"].(string)
		batch.BatchHeader.L1TxIndex = uint32(mBatchHeader["l1_tx_index"].(int64))
		batch.BatchHeader.ChainId = uint32(mBatchHeader["chain_id"].(int64))
		batch.BatchHeader.L2Blocks = make([]*sequencerv2types.BlockHeader, 0)
		for _, l2Block := range mBatchHeader["l2_blocks"].(primitive.A) {
			mL2Block := l2Block.(bson.M)
			batch.BatchHeader.L2Blocks = append(batch.BatchHeader.L2Blocks, &sequencerv2types.BlockHeader{
				BlockNumber: uint64(mL2Block["block_number"].(int64)),
				BlockHash:   mL2Block["block_hash"].(string),
			})
		}
	}

	if committeeHeader, ok := m["committee_header"]; ok {
		mCommitteeHeader := committeeHeader.(bson.M)
		batch.CommitteeHeader = &sequencerv2types.CommitteeHeader{}
		batch.CommitteeHeader.CurrentCommittee = mCommitteeHeader["current_committee"].(string)
		batch.CommitteeHeader.NextCommittee = mCommitteeHeader["next_committee"].(string)
		batch.CommitteeHeader.TotalVotingPower = uint64(mCommitteeHeader["total_voting_power"].(int64))
	}

	batch.AggSignature = m["agg_signature"].(string)
	if len(batch.AggSignature) > 0 {
		batch.PubKeys = make([]string, 0)
		for _, pubKey := range m["pub_keys"].(primitive.A) {
			batch.PubKeys = append(batch.PubKeys, pubKey.(string))
		}
	}

	batch.SequencedTime = m["sequenced_time"].(string)
	batch.FinalizedTime = m["finalized_time"].(string)

	return batch
}
