package config

// DefaultValues is the default configuration
const DefaultValues = `
[Node]
Port = "8080"
Nickname = ""
Room = "rinkeby"
PeerAddr = ""
StakingEndpoint = "https://34.229.73.193:8545"
StakingWS = "wss://mainnet.infura.io/ws/v3/f873861ee0954155b3a560eba6151d96"
AttestEndpoint = "https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79,https://mainnet.infura.io/ws/v3/f873861ee0954155b3a560eba6151d96"
Keystore = ""
StakerAddress = ""
LevelDBPath = "./leveldb"
LogLevel = 5

[Server]
GRPCPort = "9090"
PublicKey = ""

[Client]
GrpcURL = "localhost:9090"
PrivateKey = ""
PullInterval = "2s"

[Store]
BackendType = "memdb"
`
