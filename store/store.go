package store

// NewStorage returns a new database based on the given configuration.
func NewStorage(cfg Config) (Storage, error) {
	if cfg.BackendType == "memdb" {
		return NewMemDB()
	}

	return nil, nil
}
