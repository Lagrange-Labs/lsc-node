package store

import (
	"context"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/network/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB is a storage layer that interacts with MongoDB
type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoDB initializes a new MongoDB instance with the given configuration
func NewMongoDB(ctx context.Context, cfg Config) (*MongoDB, error) {
	// Set MongoDB client options
	clientOpts := options.Client().ApplyURI(cfg.URI)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	// Get a handle to the specified database
	database := client.Database(cfg.DatabaseName)
	return &MongoDB{
		client:   client,
		database: database,
	}, nil
}

// GetBlock retrieves a block by block number from the MongoDB
func (mdb *MongoDB) GetBlock(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	// Create a context with timeout for the query.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Define the filter for the MongoDB query.
	filter := bson.M{"header.block_number": blockNumber}
	// Execute the query and decode the result.
	var block types.Block
	mdb.collection.Database().Collection("blocks")
	err := mdb.collection.FindOne(ctx, filter).Decode(&block)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("block not found for block number: %d", blockNumber)
		}
		return nil, fmt.Errorf("error fetching block from MongoDB: %v", err)
	}
	return &block, nil
}

// AddBlock inserts a new block into the MongoDB
func (mdb *MongoDB) AddBlock(ctx context.Context, block *types.Block) error {
	// Create a context with timeout for the insert operation
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Insert the block into the blocks collection
	mdb.collection.Database().Collection("blocks")
	_, err := mdb.collection.InsertOne(ctx, block)
	if err != nil {
		return fmt.Errorf("error inserting block into MongoDB: %v", err)
	}
	return nil
}

// AddNode inserts a new node into the MongoDB
func (mdb *MongoDB) AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error {
	// Create a context with timeout for the insert operation
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Convert pubKey from string to *bls.PublicKey
	pk := new(bls.PublicKey)
	err := pk.Deserialize(common.FromHex(pubKey))
	if err != nil {
		return err
	}
	if err != nil {
		return fmt.Errorf("error converting public key from hex string: %v", err)
	}
	// Create a new ClientNode
	node := &network.ClientNode{
		PublicKey:    pk,
		IPAddress:    ipAdr,
		StakeAddress: stakeAdr,
	}
	// Insert the node into the nodes collection
	mdb.collection.Database().Collection("nodes")
	_, err = mdb.collection.InsertOne(ctx, node)
	if err != nil {
		return fmt.Errorf("error inserting node into MongoDB: %v", err)
	}

	return nil
}

// GetNode retrieves a node by IP address from the MongoDB
func (mdb *MongoDB) GetNode(ctx context.Context, ip string) (*network.ClientNode, error) {
	// Create a context with timeout for the query
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Define the filter for the MongoDB query.
	filter := bson.M{"ip_address": ip}
	// Execute the query and decode the result
	var node network.ClientNode
	mdb.collection.Database().Collection("nodes")
	err := mdb.collection.FindOne(ctx, filter).Decode(&node)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("node not found for IP address: %s", ip)
		}
		return nil, fmt.Errorf("error fetching node from MongoDB: %v", err)
	}

	return &node, nil
}

// GetNodeCount retrieves the number of nodes in the network from the MongoDB
func (mdb *MongoDB) GetNodeCount(ctx context.Context) (uint16, error) {
	// Create a context with timeout for the query
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Define an empty filter to count all documents in the collection.
	filter := bson.M{}
	// Execute the query and get the document count
	mdb.collection.Database().Collection("nodes")
	// Execute the query and get the document count.
	count, err := mdb.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("error counting nodes in MongoDB: %v", err)
	}

	// Convert the count to uint16
	nodeCount := uint16(count)

	return nodeCount, nil
}

// GetLastBlock retrieves the last block submitted to the network from the MongoDB.
func (mdb *MongoDB) GetLastBlock(ctx context.Context) (*types.Block, error) {
	// Create a context with timeout for the query.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Define an empty filter to query all documents in the collection.
	filter := bson.M{}

	// Define a sort option to sort documents by block number in descending order.
	opts := options.FindOne().SetSort(bson.D{{Key: "header.block_number", Value: -1}})

	// Execute the query and decode the result.
	var lastBlock types.Block
	mdb.collection.Database().Collection("blocks")
	err := mdb.collection.FindOne(ctx, filter, opts).Decode(&lastBlock)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no blocks found in the MongoDB")
		}
		return nil, fmt.Errorf("error fetching last block from MongoDB: %v", err)
	}

	return &lastBlock, nil
}
