package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func up_0001(client *mongo.Client) error {
	db := client.Database("state")
	// create blocks, nodes collections
	_ = db.Collection("blocks")
	_ = db.Collection("nodes")
	_ = db.Collection("evidences")
	return nil
}

func down_0001(client *mongo.Client) error {
	db := client.Database("state")
	// drop blocks, nodes collections
	_ = db.Collection("blocks").Drop(context.Background())
	_ = db.Collection("nodes").Drop(context.Background())
	_ = db.Collection("evidences").Drop(context.Background())
	return nil
}
