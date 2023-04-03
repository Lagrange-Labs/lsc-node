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
PublicKey = "0xd3c6b5a2698e219e9fa9b3d3f4753d0668ecf43189d57c0d14798a68f8a32c9c9be1b020eb68d1010f83eb0c1b44b14c"

[Client]
GrpcURL = "localhost:9090"
PrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
StakeAddress = "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
PullInterval = "2s"

[Store]
BackendType = "memdb"

[Sequencer]
EthereumURL = "https://34.229.73.193:8545"
StakingSCAddress = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
StackingCheckInterval = 20
`
