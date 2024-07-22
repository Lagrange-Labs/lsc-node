package migrations

import (
	"context"
	"fmt"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	convert "github.com/Lagrange-Labs/lagrange-node/store/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func up_0003(client *mongo.Client) error {
	db := client.Database("state")
	chainId := uint64(10)
	batchSize := int64(2)
	startBatchNumber := int64(0) // for testing purposes
	maxBatchNumber := int64(2) // for testing purposes

	logger.Info("Starting to update l2 blocks range for batches")
	for batchNumber := startBatchNumber; batchNumber < maxBatchNumber; batchNumber+=batchSize {
		endBatchNumber := batchNumber + batchSize
		if endBatchNumber > maxBatchNumber {
			endBatchNumber = maxBatchNumber
		}
		logger.Infof("Fetching batches from %d to %d", batchNumber, endBatchNumber)
		filter := bson.M{
			"batch_header.chain_id": chainId,
			"batch_header.batch_number": bson.M{
				"$gte": batchNumber,
				"$lt":  endBatchNumber,
			},
			"batch_header.l2_from_block_number": bson.M{"$exists": false},
		}
		cursor, err := db.Collection("batches").Find(context.Background(), filter)
		if err != nil {
			logger.Errorf("Error fetching batches from %d to %d: %v", batchNumber, endBatchNumber, err)
			return err
		}
		defer cursor.Close(context.Background())

		var batches []*sequencertypes.Batch
		for cursor.Next(context.Background()) {
			mbatch := bson.M{}
			if err := cursor.Decode(&mbatch); err != nil {
				return fmt.Errorf("error decoding batch %v", err)
			}
			batch := convert.ConvertMongoToBatch(mbatch)
			batches = append(batches, batch)
		}

		for _, batch := range batches {
			logger.Infof("Processing batch %d", batch.BatchHeader.BatchNumber)
			if len(batch.BatchHeader.L2Blocks) > 0 {
				fromL2BlockNumber := batch.BatchHeader.L2Blocks[0].BlockNumber
				toL2BlockNumber := batch.BatchHeader.L2Blocks[len(batch.BatchHeader.L2Blocks)-1].BlockNumber
				update := bson.M{
					"$set": bson.M{
						"batch_header.l2_from_block_number": fromL2BlockNumber,
						"batch_header.l2_to_block_number":   toL2BlockNumber,
					},
				}
				_, err := db.Collection("batches").UpdateOne(context.Background(), bson.M{
					"batch_header.chain_id":     batch.BatchHeader.ChainId,
					"batch_header.batch_number": batch.BatchHeader.BatchNumber,
				}, update)
				if err != nil {
					logger.Errorf("Error updating batch %d: %v", batch.BatchHeader.BatchNumber, err)
					return err
				}
				logger.Infof("Updated batch %d with fromL2BlockNumber %d and toL2BlockNumber %d",
					batch.BatchHeader.BatchNumber, fromL2BlockNumber, toL2BlockNumber)
			}
		}	 
	}
	return nil
}

func down_0003(client *mongo.Client) error {
	db := client.Database("state")
	batchesCollection := db.Collection("batches")

	fromBatchNumber := int64(0) // for testing purposes
	toBatchNumber := int64(2) // for testing purposes
	chainId := 10

	filter := bson.M{
		"batch_header.batch_number": bson.M{
			"$gte": fromBatchNumber,
			"$lte": toBatchNumber,
		},
		"batch_header.chain_id": chainId,
	}
	res, err := batchesCollection.UpdateMany(
		context.Background(),
		filter,
		bson.M{"$unset": bson.M{"l2_from_block_number": "", "l2_to_block_number": ""}},
	)
	if err != nil {
		return err
	}
	logger.Infof("Removed l2_from_block_number and l2_to_block_number fields from %d documents", res.ModifiedCount)
	return nil
}