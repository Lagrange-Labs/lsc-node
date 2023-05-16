package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
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

// AddNode adds a client node to the network.
func (db *MongoDB) AddNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	return nil
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

// GetBlock returns the block for the given block number.
func (db *MongoDB) GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error) {
	collection := db.client.Database("state").Collection("blocks")
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"chain_header.block_number": blockNumber}).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrBlockNotFound
	}
	if err != nil {
		return nil, err
	}
	return ConvertMongoToBlock(block), nil
}

// GetLastBlock returns the last block that was submitted to the network.
func (db *MongoDB) GetLastBlock(ctx context.Context) (*sequencertypes.Block, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"agg_signature": bson.M{"$ne": nil}}, sortOptions).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return nil, types.ErrBlockNotFound
	}
	if err != nil {
		return nil, err
	}
	return ConvertMongoToBlock(block), nil
}

// GetLastBlockNumber returns the last block number that was submitted to the network.
func (db *MongoDB) GetLastBlockNumber(ctx context.Context, chainID int32) (uint64, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
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

// GetLastFinalizedBlockNumber returns the last block number that was finalized.
func (db *MongoDB) GetLastFinalizedBlockNumber(ctx context.Context, chainID int32) (uint64, error) {
	collection := db.client.Database("state").Collection("blocks")
	sortOptions := options.FindOne().SetSort(bson.D{{"chain_header.block_number", -1}}) //nolint:govet
	block := bson.M{}
	err := collection.FindOne(ctx, bson.M{"agg_signature": bson.M{"$ne": nil}}, sortOptions).Decode(&block)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	chainHeader := block["chain_header"].(bson.M)
	return uint64(chainHeader["block_number"].(int64)), nil
}

// GetNode returns the node for the given IP address.
func (db *MongoDB) GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error) {
	return nil, nil
}

// GetNodeCount returns the number of nodes in the network.
func (db *MongoDB) GetNodeCount(ctx context.Context) (uint16, error) {
	return 0, nil
}

// GetNodesByStatuses returns the nodes with the given statuses.
func (db *MongoDB) GetNodesByStatuses(ctx context.Context, statuses []sequencertypes.NodeStatus) ([]sequencertypes.ClientNode, error) {
	return nil, nil
}

// UpdateBlock updates the block in the database.
func (db *MongoDB) UpdateBlock(ctx context.Context, block *sequencertypes.Block) error {
	return nil
}

// UpdateNode updates the node status in the database.
func (db *MongoDB) UpdateNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	return nil
}
