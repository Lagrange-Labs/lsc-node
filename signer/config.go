package signer

type KMSType int

// KMSConfig is the configuration for the KMS provider.
type ProviderConfig struct {
	Type string `mapstructure:"Type"`
	// Local
	LocalConfig *LocalConfig `mapstructure:"LocalConfig"`
	// AWS KMS
	AWSKMSConfig *AWSKMSConfig `mapstructure:"AWSKMSConfig"`
}

// LocalConfig is the configuration for the local provider.
type LocalConfig struct {
	PrivateKeyPath  string `mapstructure:"PrivateKeyPath"`
	PasswordKeyPath string `mapstructure:"PasswordKeyPath"`
}

// AWSKMSConfig is the configuration for the AWS KMS provider.
type AWSKMSConfig struct {
	Region string `mapstructure:"Region"`
	KeyID  string `mapstructure:"KeyID"`
}

// Config is the configuration for the signer.
type Config struct {
	ProviderConfigs []ProviderConfig `mapstructure:"ProviderConfigs"`
	CertKeyPath     string           `mapstructure:"CertKeyPath"`
}
