package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"

[Client]
GrpcURL = "localhost:9090"
PrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
StakeAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
PullInterval = "2s"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://localhost:27017"

[Sequencer]
Chain = "arbitrum"
FromBlockNumber = 1

[Governance]
EthereumURL = "https://34.229.73.193:8545"
StakingSCAddress = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
StakingCheckInterval = 20

[Consensus]
ProposerPrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
RoundInterval = "5s"
RoundLimit = "30s"
`
