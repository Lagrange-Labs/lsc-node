package migrations

import (
	"context"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func up_0002(client *mongo.Client) error {
	db := client.Database("state")
	batchesCollection := db.Collection("batches")
	batchSize := 2            // for testing purposes
	chainId := uint64(421614) // keeping chainId static to avoid migration issues
	lastBatchNumber := int64(0)
	maxBatchNumber := int64(2) // for testing purposes

	// cache committee details for a range of batches to reduce the number of queries
	var cachedCommitteeDetails map[string]interface{}
	logger.Info("Updating operators field in batches")
	for {
		logger.Infof("Fetching batches from %d to %d", lastBatchNumber, lastBatchNumber+int64(batchSize))
		filter := bson.M{
			"batch_header.batch_number": bson.M{
				"$gt":  lastBatchNumber,
				"$lte": maxBatchNumber,
			},
			"batch_header.chain_id": chainId,
			"agg_signature":         bson.M{"$ne": ""},
			"operators":             bson.M{"$exists": false},
		}
		// Fetch the batch of documents from the database
		cursor, err := batchesCollection.Find(context.Background(), filter, options.Find().SetLimit(int64(batchSize)).SetSort(bson.M{"batch_header.batch_number": 1}))
		if err != nil {
			logger.Infof("Error fetching batches: %v", err)
			return err
		}

		defer cursor.Close(context.Background())

		var batchesProcessed bool
		for cursor.Next(context.Background()) {
			batchesProcessed = true
			var batch map[string]interface{}
			if err := cursor.Decode(&batch); err != nil {
				logger.Infof("Error decoding batch: %v", err)
				return err
			}
			batchNumber := batch["batch_header"].(map[string]interface{})["batch_number"].(int64)
			current_committee_root := batch["committee_header"].(map[string]interface{})["current_committee"].(string)
			logger.Infof("Processing batch %d", batchNumber)

			// If the cache is not valid for the current batch, update the cache
			if cachedCommitteeDetails == nil || current_committee_root != cachedCommitteeDetails["current_committee_root"].(string) {
				logger.Infof("Updating cache for batch=%d and committee root=%s", batchNumber, current_committee_root)
				// Get the committee details for the batch
				committee_details, err := getCommitteeDetailsForBatch(db, current_committee_root, chainId)
				if err != nil {
					logger.Infof("Error getting committee details for batch: %v", err)
					return err
				}
				cachedCommitteeDetails = committee_details
			}

			operators := make([]string, len(batch["pub_keys"].(primitive.A)))
			for i, pubkey := range batch["pub_keys"].(primitive.A) {
				pubkeyStr, ok := pubkey.(string)
				if !ok {
					logger.Infof("Error parsing pubkey for batch %d", batchNumber)
					return fmt.Errorf("error parsing pubkey for batch %d", batchNumber)
				}
				for _, operator := range cachedCommitteeDetails["operators"].(primitive.A) {
					operatorMap, ok := operator.(map[string]interface{})
					if !ok {
						logger.Infof("Error parsing operator for batch %d", batchNumber)
						return fmt.Errorf("error parsing operator for batch %d", batchNumber)
					}
					if operatorMap["public_key"].(string) == pubkeyStr {
						operators[i] = operatorMap["stake_address"].(string)
						break
					}
				}
			}

			// Update the batch with the operators using batch_number and chain_id
			_, err = batchesCollection.UpdateOne(
				context.Background(),
				bson.M{
					"batch_header.batch_number": batchNumber,
					"batch_header.chain_id":     chainId,
				},
				bson.M{"$set": bson.M{"operators": operators}},
			)
			if err != nil {
				logger.Infof("Error updating batch: %v", err)
				return err
			}
			logger.Info("Batch processed successfully.")
			lastBatchNumber = batchNumber
		}
		if !batchesProcessed {
			logger.Info("No batches processed")
			break
		}
	}
	return nil
}

func down_0002(client *mongo.Client) error {
	db := client.Database("state")
	batchesCollection := db.Collection("batches")

	fromBatchNumber := 0  // for testing purposes
	toBatchNumber := 3700 // for testing purposes
	chainId := 42161

	filter := bson.M{
		"batch_header.batch_number": bson.M{
			"$gte": fromBatchNumber,
			"$lte": toBatchNumber,
		},
		"batch_header.chain_id": chainId,
	}

	// Remove the operators field from matched documents
	res, err := batchesCollection.UpdateMany(
		context.Background(),
		filter,
		bson.M{"$unset": bson.M{"operators": ""}},
	)
	if err != nil {
		return err
	}
	logger.Infof("Removed operators field from %d documents", res.ModifiedCount)
	return nil
}

// getCommitteeDetailsForBatch fetches the committee details for the given batch number and chain id
func getCommitteeDetailsForBatch(db *mongo.Database, committeeRoot string, chainId uint64) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(1)
	cursor, err := db.Collection("committee_roots").Find(ctx, bson.M{"chain_id": chainId, "current_committee_root": committeeRoot}, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("committee root not found")
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	var committeeRootDetails map[string]interface{}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&committeeRootDetails); err != nil {
			return nil, err
		}
		return committeeRootDetails, nil
	}

	return nil, fmt.Errorf("committee root not found")
}
