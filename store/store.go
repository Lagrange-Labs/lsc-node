package store

import (
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
	"github.com/Lagrange-Labs/lagrange-node/store/mongodb"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
)

// NewStorage returns a new database based on the given configuration.
func NewStorage(cfg *Config) (types.Storage, error) {
	if cfg.BackendType == "memdb" {
		return memdb.NewMemDB()
	} else if cfg.BackendType == "mongodb" {
		return mongodb.NewMongoDB(cfg.DBPath)
	}

	return nil, nil
}
