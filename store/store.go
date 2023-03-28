package store

// NewDB returns a new database based on the given configuration.
func NewDB(cfg Config) (DB, error) {
	if cfg.BackendType == "memdb" {
		return NewMemDB()
	}

	return nil, nil
}
