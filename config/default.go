package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"
PublicKey = "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4"

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
StakingCheckInterval = 20
`
