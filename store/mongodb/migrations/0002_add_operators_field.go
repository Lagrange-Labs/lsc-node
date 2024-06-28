package migrations

import (
	"context"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func up_0002(client *mongo.Client) error {
	db := client.Database("state")
	batchesCollection := db.Collection("batches")
	batchSize := 2 // for testing purposes
	chainId := uint64(421614) // keeping chainId static to avoid migration issues
	lastBatchNumber := int64(0)
	maxBatchNumber := int64(2) // for testing purposes
	
	// cache committee details for a range of batches to reduce the number of queries
	var cachedCommitteeDetails map[string]interface{}
	var cacheValidForRangeEnd int64
	logger.Info("Updating operators field in batches")
	for {
		logger.Infof("Fetching batches from %d to %d", lastBatchNumber, lastBatchNumber+int64(batchSize))
		filter := bson.M{"batch_header.batch_number": bson.M{"$gt": lastBatchNumber, "$lte": maxBatchNumber}, "batch_header.chain_id": chainId, "agg_signature": bson.M{"$ne": ""}, "operators": bson.M{"$exists": false}}
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
			logger.Infof("Processing batch %d", batchNumber)

			// If the cache is not valid for the current batch, update the cache
			if cachedCommitteeDetails == nil || batchNumber > cacheValidForRangeEnd {
				logger.Infof("Updating cache for batch %d", batchNumber)
				// Get the committee details for the batch
				committee_details, err := getCommitteeDetailsForBatch(batchesCollection, batchNumber, chainId)
				if err != nil {
					logger.Infof("Error getting committee details for batch: %v", err)
					return err
				}
				cachedCommitteeDetails = committee_details
				if committeeDetailsSlice, ok := committee_details["committee_details"].(primitive.A); ok {
					if len(committeeDetailsSlice) > 0 {
						if committeeDetail, ok := committeeDetailsSlice[0].(map[string]interface{}); ok {
							cacheValidForRangeEnd = committeeDetail["epoch_end_block_number"].(int64)
						} else {
							logger.Infof("Error parsing committee details for batch %d", batchNumber)
							return fmt.Errorf("error parsing committee details for batch %d", batchNumber)
						}
					} else {
						logger.Infof("No committee details found for batch %d", batchNumber)
						continue
					}
				}
			}
			
			operators := make([]string, len(batch["pub_keys"].(primitive.A)))
			for i, pubkey := range batch["pub_keys"].(primitive.A) {
				pubkeyStr, ok := pubkey.(string)
				if !ok {
					logger.Infof("Error parsing pubkey for batch %d", batchNumber)
					return fmt.Errorf("error parsing pubkey for batch %d", batchNumber)
				}
				for _, operator := range cachedCommitteeDetails["committee_details"].(primitive.A)[0].(map[string]interface{})["operators"].(primitive.A) {
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

			logger.Infof("Operators to update: %v", operators)

			// Update the batch with the operators using batch_number and chain_id
			_, err = batchesCollection.UpdateOne(
				context.Background(),
				bson.M{
					"batch_header.batch_number": batchNumber,
					"batch_header.chain_id": chainId,
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

	fromBatchNumber := 1 // for testing purposes
	toBatchNumber := 2 // for testing purposes
	chainId := 421614

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
func getCommitteeDetailsForBatch(batchesCollection *mongo.Collection, batchNumber int64, chainId uint64) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "batch_header.batch_number", Value: batchNumber},
			{Key: "batch_header.chain_id", Value: chainId},
		}}},

		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "committee_roots"},
			{Key: "let", Value: bson.D{
				{Key: "current_committee", Value: "$committee_header.current_committee"},
				{Key: "chain_id", Value: "$batch_header.chain_id"},
				{Key: "l1_block_number", Value: "$batch_header.l1_block_number"},
			}},
			{Key: "pipeline", Value: mongo.Pipeline{
				{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "$eq", Value: bson.A{"$current_committee_root", "$$current_committee"}}},
							bson.D{{Key: "$eq", Value: bson.A{"$chain_id", "$$chain_id"}}},
							bson.D{{Key: "$lte", Value: bson.A{"$epoch_start_block_number", "$$l1_block_number"}}},
							bson.D{{Key: "$gte", Value: bson.A{"$epoch_end_block_number", "$$l1_block_number"}}},
						}},
					}},
				}}},
			}},
			{Key: "as", Value: "committee_details"},
		}}},

		{{Key: "$project", Value: bson.D{
			{Key: "agg_signature", Value: 0},
			{Key: "batch_header", Value: 0},
			{Key: "finalized_time", Value: 0},
			{Key: "sequenced_time", Value: 0},
			{Key: "proposer_pub_key", Value: 0},
			{Key: "proposer_signature", Value: 0},
		}}},
	}

	cursor, err := batchesCollection.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Errorf("Error executing aggregation pipeline: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		logger.Errorf("Error parsing aggregation result: %v", err)
		return nil, err
	}

	if len(results) == 0 {
		logger.Infof("No committee root details found for batchNumber %d and chainId %d", batchNumber, chainId)
		return nil, fmt.Errorf("no committee root details found")
	}

	return results[0], nil
}