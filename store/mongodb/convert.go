package mongodb

import (
	"fmt"
	"reflect"
	"strings"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
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
