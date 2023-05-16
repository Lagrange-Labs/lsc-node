package migrations

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Migration is a single migration.
type Migration struct {
	Step        uint32    `bson:"step"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	up          func(client *mongo.Client) error
	down        func(client *mongo.Client) error
}

// MigrationManager is the manager for all migrations.
type MigrationManager struct {
	migrateCollection *mongo.Collection
	client            *mongo.Client
	migrations        []*Migration
}

var m = &MigrationManager{}

// RegisterDB registers the database to the migration manager.
func RegisterDB(uri string) *MigrationManager {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db := client.Database("migrate")
	m.migrateCollection = db.Collection("migrations")
	m.client = client
	return m
}

// RegisterMigration registers a migration to the migration manager.
func RegisterMigration(name string, up, down func(client *mongo.Client) error) error {
	slice := strings.Split(name, "_")
	num, err := strconv.ParseUint(slice[0], 10, 32)
	if err != nil {
		return err
	}
	m.migrations = append(m.migrations, &Migration{
		Step:        uint32(num),
		Description: slice[1],
		up:          up,
		down:        down,
	})

	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Step < m.migrations[j].Step
	})

	return nil
}

// MigrateUp migrates the database up to the given step.
func (mm *MigrationManager) MigrateUp(step uint32) error {
	if step == 0 {
		step = mm.migrations[len(mm.migrations)-1].Step
	}
	if mm.migrations[len(mm.migrations)-1].Step < step {
		return fmt.Errorf("cannot migrate up, the last migratable step is %d", mm.migrations[len(mm.migrations)-1].Step)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var migration Migration
	if err := mm.migrateCollection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"step": -1})).Decode(&migration); err != nil {
		if err == mongo.ErrNoDocuments {
			migration = Migration{Step: 0}
		} else {
			return err
		}
	}

	lastStep := migration.Step
	if step <= lastStep {
		return fmt.Errorf("cannot migrate up, the recent migrated step is %d", lastStep)
	}

	for i := 0; i < len(mm.migrations); i++ {
		if mm.migrations[i].Step > step {
			break
		}
		if mm.migrations[i].Step <= lastStep {
			continue
		}
		if err := mm.migrations[i].up(mm.client); err != nil {
			return err
		}
		mm.migrations[i].CreatedAt = time.Now()
		if _, err := mm.migrateCollection.InsertOne(ctx, mm.migrations[i]); err != nil {
			return err
		}
	}
	return nil
}

// MigrateDown migrates the database down to the given step.
func (mm *MigrationManager) MigrateDown(step uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var migration Migration
	if err := mm.migrateCollection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"step": -1})).Decode(&migration); err != nil {
		if err == mongo.ErrNoDocuments {
			migration = Migration{Step: 0}
		} else {
			return err
		}
	}
	lastStep := migration.Step
	if step >= lastStep {
		return fmt.Errorf("cannot migrate down, the recent migrated step is %d", lastStep)
	}

	for i := len(mm.migrations) - 1; i >= 0; i-- {
		if mm.migrations[i].Step <= step {
			break
		}
		if mm.migrations[i].Step > lastStep {
			continue
		}
		if err := mm.migrations[i].down(mm.client); err != nil {
			return err
		}
		if _, err := mm.migrateCollection.DeleteOne(ctx, bson.M{"step": mm.migrations[i].Step}); err != nil {
			return err
		}
	}
	return nil
}
