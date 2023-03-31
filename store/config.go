package store

// Config is the configuration for the store.
type Config struct {
	// BackendType is the database type
	BackendType string `mapstructure:"BackendType"`
	// URI is the database connection string
	URI string `mapstructure:"URI" default:"mongodb://mongo:27017"`
	// DatabaseName is the name of the database
	DatabaseName string `mapstructure:"DatabaseName" default:"lagrangeDB"`
}
