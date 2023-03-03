package node

type Config struct {
	// Server listening port
	Port string `mapstructure:"Port"`
	// Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later
	Nickname string `mapstructure:"Nickname"`
	// Room / Network
	Room string `mapstructure:"Room"`
	// Remote Peer Address
	PeerAddr string `mapstructure:"PeerAddr"`
	// Staking Endpoint URL:Port
	StakingEndpoint string `mapstructure:"StakingEndpoint"`
	// Staking Listening Endpoint wss://URL/ws
	StakingWS string `mapstructure:"StakingWS"`
	// Attestation Endpoint (URL:Port).  Multiple endpoints can be specified with a comma delimiter.
	AttestEndpoint string `mapstructure:"AttestEndpoint"`
	// Keystore Path
	Keystore string `mapstructure:"Keystore"`
	// Staker Address
	StakerAddress string `mapstructure:"StakerAddress"`
	// Level DB storage location
	LevelDBPath string `mapstructure:"LevelDBPath"`
	// Logging output level (1:INFO, 2:NOTICE, 3:WARNING, 4:ERROR, 5:DEBUG)
	LogLevel int `mapstructure:"LogLevel"`
}
