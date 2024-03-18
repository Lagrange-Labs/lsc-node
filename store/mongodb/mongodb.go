package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	govtypes "github.com/Lagrange-Labs/lagrange-node/governance/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ types.Storage = (*MongoDB)(nil)

type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB creates a new MongoDB instance.
func NewMongoDB(uri string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(200)
	clientOptions.SetMinPoolSize(5)
	clientOptions.SetMaxConnIdleTime(10 * time.Minute)

	client, err := mongo.Connect(context.Background(), clientOptions)
	return &MongoDB{
		client: client,
	}, err
}

// AddNode adds a joined client node to the network.
func (db *MongoDB) AddNode(ctx context.Context, node *networktypes.ClientNode) error {
	collection := db.client.Database("state").Collection("nodes")
	_, err := collection.UpdateOne(ctx, bson.M{"stake_address": node.StakeAddress, "chain_id": node.ChainID}, bson.M{"$set": node}, options.Update().SetUpsert(true))
	return err
}

// AddBlock adds a block to the storage.c
func (db *MongoDB) AddBlock(ctx context.Context, block *sequencertypes.Block) error {
	mBlock, err := ConvertProtobufToMongo(block)
	if err != nil {
		return fmt.Errorf("failed to convert block to mongo: %w", err)
	}

	_, err = db.client.Database("state").Collection("blocks").InsertOne(ctx, mBlock)
	return err
}

// UpdateBlock updates the block in the database.
func (db *MongoDB) UpdateBlock(ctx context.Context, block *sequencertypes.Block) error {
	mBlock, err := ConvertProtobufToMongo(block)
	if err != nil {
		return fmt.Errorf("failed to convert block to mongo: %w", err)
	}

	collection := db.client.Database("state").Collection("blocks")
	_, err = collection.UpdateOne(ctx, bson.M{"chain_header.block_number": block.BlockNumber(), "chain_header.chain_id": block.ChainHeader.ChainId}, bson.M{"$set": mBlock})
	return err
}

// GetBlock returns the block for the given block number.
func (db *MongoDB) GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error) {
	collection := db.client.Database("state").Collection("blocks")
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"chain_header.block_number": blockNumber, "chain_header.chain_id": chainID}).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrBlockNotFound
	}
	if err != nil {
		return nil, err
	}
	return ConvertMongoToBlock(block), nil
}

// GetBlocks returns the `count` blocks starting from `fromBlockNumber`.
func (db *MongoDB) GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error) {
	collection := db.client.Database("state").Collection("blocks")
	filter := bson.M{"chain_header.block_number": bson.M{"$gte": fromBlockNumber, "$lt": fromBlockNumber + uint64(count)}, "chain_header.chain_id": chainID}
	cursor, err := collection.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrBlockNotFound
	}
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	blocks := make([]*sequencertypes.Block, 0)
	for cursor.Next(ctx) {
		block := bson.M{}
		err := cursor.Decode(&block)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, ConvertMongoToBlock(block))
	}
	return blocks, nil
}

// GetLastFinalizedBlock returns the last finalized block for the given chainID.
func (db *MongoDB) GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"pub_keys": bson.M{"$ne": nil}, "chain_header.chain_id": chainID}, sortOptions).Decode(&block)
	if err == mongo.ErrNoDocuments {
		sortOptions = options.FindOne().SetSort(bson.D{{"chain_header.block_number", 1}}) //nolint:govet
		err = collection.FindOne(ctx, bson.M{"chain_header.chain_id": chainID}, sortOptions).Decode(&block)
		if err == mongo.ErrNoDocuments {
			return nil, types.ErrBlockNotFound
		}
	}
	return ConvertMongoToBlock(block), err
}

// GetLastFinalizedBlockNumber returns the last block number that was stored.
func (db *MongoDB) GetLastFinalizedBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}).SetProjection(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"pub_keys": bson.M{"$ne": nil}, "chain_header.chain_id": chainID}, sortOptions).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	chainHeader := block["chain_header"].(bson.M)
	return uint64(chainHeader["block_number"].(int64)), nil
}

// GetLastBlockNumber returns the last block number that was stored.
func (db *MongoDB) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}).SetProjection(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"chain_header.chain_id": chainID}, sortOptions).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	chainHeader := block["chain_header"].(bson.M)
	return uint64(chainHeader["block_number"].(int64)), nil
}

// AddBatch adds a new batch to the database.
func (db *MongoDB) AddBatch(ctx context.Context, batch *sequencerv2types.Batch) error {
	mBatch, err := ConvertProtobufToMongo(batch)
	if err != nil {
		return fmt.Errorf("failed to convert batch to mongo: %w", err)
	}

	collection := db.client.Database("state").Collection("batches")
	_, err = collection.InsertOne(ctx, mBatch)
	return err
}

// UpdateBatch updates the batch in the database.
func (db *MongoDB) UpdateBatch(ctx context.Context, batch *sequencerv2types.Batch) error {
	mBatch, err := ConvertProtobufToMongo(batch)
	if err != nil {
		return fmt.Errorf("failed to convert batch to mongo: %w", err)
	}

	collection := db.client.Database("state").Collection("batches")
	_, err = collection.UpdateOne(ctx, bson.M{"batch_header.batch_number": batch.BatchNumber(), "batch_header.chain_id": batch.ChainID()}, bson.M{"$set": mBatch})
	return err
}

// GetLastFinalizedBatchNumber returns the last finalized batch number for the given chainID.
func (db *MongoDB) GetLastFinalizedBatchNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("batches")
	sortOptions := options.FindOne().SetSort(bson.D{{"batch_header.batch_number", -1}}).SetProjection(bson.D{{"batch_header.batch_number", 1}}) //nolint:govet
	batch := bson.M{}
	err := collection.FindOne(ctx, bson.M{"pub_keys": bson.M{"$ne": nil}, "batch_header.chain_id": chainID}, sortOptions).Decode(&batch)
	if err == mongo.ErrNoDocuments {
		sortOptions = options.FindOne().SetSort(bson.D{{"batch_header.batch_number", 1}}).SetProjection(bson.D{{"batch_header.batch_number", 1}}) //nolint:govet
		err = collection.FindOne(ctx, bson.M{"batch_header.chain_id": chainID}, sortOptions).Decode(&batch)
		if err == mongo.ErrNoDocuments {
			return 0, types.ErrBatchNotFound
		} else if err != nil {
			return 0, err
		}
		return uint64(batch["batch_header"].(bson.M)["batch_number"].(int64) - 1), nil
	}

	return uint64(batch["batch_header"].(bson.M)["batch_number"].(int64)), err
}

// GetLastBatchNumber returns the last batch number that was stored to the db.
func (db *MongoDB) GetLastBatchNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("batches")
	sortOptions := options.FindOne().SetSort(bson.D{{"batch_header.batch_number", -1}}).SetProjection(bson.D{{"batch_header.batch_number", 1}}) //nolint:govet
	batch := bson.M{}
	err := collection.FindOne(ctx, bson.M{"batch_header.chain_id": chainID}, sortOptions).Decode(&batch)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	return uint64(batch["batch_header"].(bson.M)["batch_number"].(int64)), err
}

// GetBatch returns the batch for the given batch number and chainID.
func (db *MongoDB) GetBatch(ctx context.Context, chainID uint32, batchNumber uint64) (*sequencerv2types.Batch, error) {
	collection := db.client.Database("state").Collection("batches")
	batch := bson.M{}
	err := collection.FindOne(ctx, bson.M{"batch_header.batch_number": batchNumber, "batch_header.chain_id": chainID}).Decode(&batch)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrBatchNotFound
	}
	if err != nil {
		return nil, err
	}
	return ConvertMongoToBatch(batch), nil
}

// GetLastEvidenceBlockNumber returns the last block number of the submitted evidence.
func (db *MongoDB) GetLastEvidenceBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("evidences")
	sortOptions := options.FindOne().SetSort(bson.D{{"block_number", -1}}) //nolint:govet
	evidence := bson.M{}
	err := collection.FindOne(ctx, bson.M{"chain_id": chainID, "status": true}, sortOptions).Decode(&evidence)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return uint64(evidence["block_number"].(int64)), nil
}

// GetNodeByStakeAddr returns the node for the given stake address.
func (db *MongoDB) GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*networktypes.ClientNode, error) {
	collection := db.client.Database("state").Collection("nodes")
	node := networktypes.ClientNode{}
	err := collection.FindOne(ctx, bson.M{"stake_address": stakeAddress, "chain_id": chainID}).Decode(&node)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrNodeNotFound
	}
	return &node, err
}

// GetNodesByStatuses returns the nodes with the given statuses.
func (db *MongoDB) GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus, chainID uint32) ([]networktypes.ClientNode, error) {
	collection := db.client.Database("state").Collection("nodes")
	filter := bson.M{
		"status": bson.M{
			"$in": statuses,
		},
		"chain_id": chainID,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	nodes := []networktypes.ClientNode{}

	for cursor.Next(ctx) {
		node := networktypes.ClientNode{}
		err := cursor.Decode(&node)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}
	return nodes, cursor.Err()
}

// AddEvidences adds the given evidences to the database.
func (db *MongoDB) AddEvidences(ctx context.Context, evidences []*contypes.Evidence) error {
	mEvidences := []interface{}{}
	for _, evidence := range evidences {
		mEvidences = append(mEvidences, evidence)
	}
	collection := db.client.Database("state").Collection("evidences")
	_, err := collection.InsertMany(ctx, mEvidences)
	return err
}

// UpdateEvidence updates the given evidence in the database.
func (db *MongoDB) UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error {
	collection := db.client.Database("state").Collection("evidences")
	filter := bson.M{"block_number": evidence.BlockNumber, "chain_id": evidence.ChainID, "operator": evidence.Operator}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// GetEvidences returns the evidences for the given block range.
func (db *MongoDB) GetEvidences(ctx context.Context, chainID uint32, fromBlockNumber, toBlockNumber uint64, limit, offset int64) ([]*contypes.Evidence, error) {
	collection := db.client.Database("state").Collection("evidences")
	filter := bson.M{"chain_id": chainID, "block_number": bson.M{"$gte": fromBlockNumber, "$lte": toBlockNumber}}
	sortOptions := options.Find().SetSort(bson.D{{"block_number", 1}}).SetLimit(limit).SetSkip(offset) //nolint:govet
	cursor, err := collection.Find(ctx, filter, sortOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	evidences := []*contypes.Evidence{}
	for cursor.Next(ctx) {
		evidence := contypes.Evidence{}
		err := cursor.Decode(&evidence)
		if err != nil {
			return nil, err
		}
		evidences = append(evidences, &evidence)
	}
	return evidences, nil
}

// UpdateCommitteeRoot updates the committee root in the database.
func (db *MongoDB) UpdateCommitteeRoot(ctx context.Context, committeeRoot *govtypes.CommitteeRoot) error {
	collection := db.client.Database("state").Collection("committee_roots")
	filter := bson.M{"chain_id": committeeRoot.ChainID, "epoch_end_block_number": committeeRoot.EpochEndBlockNumber}
	update := bson.M{"$set": bson.M{"current_committee_root": committeeRoot.CurrentCommitteeRoot, "total_voting_power": committeeRoot.TotalVotingPower, "epoch_number": committeeRoot.EpochNumber, "operators": committeeRoot.Operators}}
	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

// GetCommitteeRoot returns the first committee root which EpochBlockNumber is greater than or equal to the given l1BlockNumber.
func (db *MongoDB) GetCommitteeRoot(ctx context.Context, chainID uint32, l1BlockNumber uint64) (*govtypes.CommitteeRoot, error) {
	collection := db.client.Database("state").Collection("committee_roots")
	filter := bson.M{"chain_id": chainID, "epoch_end_block_number": bson.M{"$gte": l1BlockNumber}, "epoch_start_block_number": bson.M{"$lte": l1BlockNumber}}
	sortOptions := options.FindOne().SetSort(bson.D{{"epoch_block_number", 1}}) //nolint:govet
	committeeRoot := govtypes.CommitteeRoot{}
	err := collection.FindOne(ctx, filter, sortOptions).Decode(&committeeRoot)
	return &committeeRoot, err
}

// GetLastCommitteeEpochNumber returns the last committee epoch number for the given chainID.
func (db *MongoDB) GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error) {
	collection := db.client.Database("state").Collection("committee_roots")
	sortOptions := options.FindOne().SetSort(bson.D{{"epoch_number", -1}}).SetProjection(bson.D{{"epoch_number", 1}}) //nolint:govet
	filter := bson.M{"chain_id": chainID}
	committeeRoot := govtypes.CommitteeRoot{}
	err := collection.FindOne(ctx, filter, sortOptions).Decode(&committeeRoot)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	return committeeRoot.EpochNumber, err
}
