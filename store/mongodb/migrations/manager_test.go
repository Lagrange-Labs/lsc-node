package migrations

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestRegisterMigration(t *testing.T) {
	up := func(client *mongo.Client) error { return nil }
	down := func(client *mongo.Client) error { return nil }

	require.NoError(t, RegisterMigration("0001_init", up, down))
	require.NoError(t, RegisterMigration("0002_createtable", up, down))
}

func TestMigrate(t *testing.T) {
	up := func(client *mongo.Client) error { return nil }
	down := func(client *mongo.Client) error { return nil }

	mm := RegisterDB("mongodb://localhost:27017")

	require.NoError(t, RegisterMigration("0001_init", up, down))
	require.NoError(t, RegisterMigration("0002_createtable", up, down))
	require.NoError(t, RegisterMigration("0003_createindex", up, down))
	require.NoError(t, RegisterMigration("0004_addfield", up, down))

	require.NoError(t, mm.MigrateUp(2))
	require.NoError(t, mm.MigrateDown(1))
	require.NoError(t, mm.MigrateUp(4))
	require.NoError(t, mm.MigrateDown(3))

	require.Error(t, mm.MigrateUp(5))
	require.Error(t, mm.MigrateUp(2))
	require.Error(t, m.MigrateDown(4))
}
