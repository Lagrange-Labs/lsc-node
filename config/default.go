package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"

[Client]
GrpcURL = "127.0.0.1:9090"
Chain = "mock"
EthereumURL = "http://localhost:8545"
CommitteeSCAddress = "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06"
BLSPrivateKey = "0x00000000000000000000000000000000000000000000000000000000499602d3"
ECDSAPrivateKey = "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c"
PullInterval = "100ms"
BLSCurve = "BN254"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://127.0.0.1:27017"

[Sequencer]
Chain = "mock"
FromL1BlockNumber = 20
FromL2BlockNumber = 20
EthereumURL = "http://localhost:8545"
CommitteeSCAddress = "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06"
StakingCheckInterval = "2s"

[RpcClient]

	[RpcClient.Optimism]
	RPCURL = "http://localhost:8545"
	L1RPCURL = "http://localhost:8545"
	BeginBlockNumber = 80
	BatchInbox = "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f"
	BatchSender = "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f"
	
	[RpcClient.Arbitrum]
	RPCURL = "http://localhost:8545"

	[RpcClient.Mantle]
	RPCURL = "http://localhost:8545"
	L1RPCURL = "http://localhost:8545"
	BatchStorageAddr = "0xbB9dDB1020F82F93e45DA0e2CFbd27756DA36956"

	[RpcClient.Mock]
	RPCURL = "http://localhost:8545"

[Consensus]
OperatorAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
ProposerPrivateKey = "0x8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b"
RoundInterval = "500ms"
RoundLimit = "30s"
BLSCurve = "BN254"
`
