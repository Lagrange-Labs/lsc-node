package utils

import (
	"flag"
)

type Opts struct {
	port            int
	nick            string
	room            string
	peerAddr        string
	stakingEndpoint string
	stakingWS       string
	attestEndpoint  string
	keystore        string
	address         string
	leveldb         string
	logLevel        int
}

func GetOpts() *Opts {
	// Parse Port
	portPtr := flag.Int("port", 8081, "Server listening port")
	// Parse Nickname
	nickPtr := flag.String("nick", "", "Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later.")
	// Parse Room
	roomPtr := flag.String("room", "rinkeby", "Room / Network")
	// Parse Remote Peer
	peerAddrPtr := flag.String("peerAddr", "", "Remote Peer Address")
	// Parse ETH (Staking) URL
	stakingEndpointPtr := flag.String("stakingEndpoint", "https://34.229.73.193:8545", "Staking Endpoint URL:Port")
	// Parse ETH (Staking) WS
	stakingWSPtr := flag.String("stakingWS", "wss://mainnet.infura.io/ws/v3/f873861ee0954155b3a560eba6151d96", "Staking Listening Endpoint wss://URL/ws")
	// Parse ETH (Attestation) URL
	attestEndpointPtr := flag.String("attestEndpoint", "https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79,https://mainnet.infura.io/ws/v3/f873861ee0954155b3a560eba6151d96", "Attestation Endpoint (URL:Port).  Multiple endpoints can be specified with a comma delimiter.")
	// Parse Keystore Path
	keystorePtr := flag.String("keystore", "", "/path/to/keystore")
	// Parse Address
	addressPtr := flag.String("address", "", "Staker Address")
	// LevelDB Location
	leveldbPtr := flag.String("leveldb", "./leveldb", "Level DB storage location")
	// Log Level
	logLevelPtr := flag.Int("loglevel", 5, "Logging output level (1:INFO, 2:NOTICE, 3:WARNING, 4:ERROR, 5:DEBUG)")

	flag.Parse()

	res := Opts{
		*portPtr,
		*nickPtr,
		*roomPtr,
		*peerAddrPtr,
		*stakingEndpointPtr,
		*stakingWSPtr,
		*attestEndpointPtr,
		*keystorePtr,
		*addressPtr,
		*leveldbPtr,
		*logLevelPtr}
	return &res
}
