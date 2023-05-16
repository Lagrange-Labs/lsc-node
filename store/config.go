package store

// Config is the configuration for the store.
type Config struct {
	// BackendType is the database type
	BackendType string `mapstructure:"BackendType"`
	// DBPath is the path to the database
	DBPath string `mapstructure:"DBPath"`
}
