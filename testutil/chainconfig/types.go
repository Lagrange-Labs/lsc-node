package chainconfig

// ChainConfig is the configuration for the chain
type ChainConfig struct {
	ChainName        string   `json:"chain_name"`
	ChainID          uint32   `json:"chain_id"`
	Operators        []string `json:"operators"`
	ECDSAPrivateKeys []string `json:"ecdsa_priv_keys"`
	BLSPublicKeys    []string `json:"bls_pub_keys"`
	BLSPrivateKeys   []string `json:"bls_priv_keys"`
}
