package store

// Config is the configuration for the store.
type Config struct {
	// BackendType is the database type
	BackendType string `mapstructure:"BackendType"`
}
