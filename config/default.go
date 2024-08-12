package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"

[Client]
GrpcURLs = "127.0.0.1:9090"
Chain = "mock"
EthereumURL = "http://localhost:8545"
OperatorAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
CommitteeSCAddress = "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06"
SignerServerURL = "" #"127.0.0.1:50051"
BLSKeyAccountID = "" #"bls-sign-key-0"
SignerKeyAccountID = "" #"ecdsa-signer-key-0"
BLSKeystorePath = "./testutil/vector/config/bls_0.json"
BLSKeystorePassword = "password_localtest"
BLSKeystorePasswordPath = ""
SignerECDSAKeystorePath = "./testutil/vector/config/ecdsa_0.json"
SignerECDSAKeystorePassword = "password_localtest"
SignerECDSAKeystorePasswordPath = ""
PullInterval = "100ms"
BLSCurve = "BN254"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://127.0.0.1:27017"

[Sequencer]
Chain = "mock"
FromL1BlockNumber = 20
EthereumURL = "http://localhost:8545"
CommitteeSCAddress = "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06"
EigenDMSCAddress = "0xF2740f6A6333c7B405aD7EfC68c74adAd83cC30D"
StakingCheckInterval = "2s"

[RpcClient]

	[RpcClient.Optimism]
	RPCURL = "http://localhost:8545"
	L1RPCURL = "http://localhost:8545"
	BeaconURL = "http://localhost:8545"
	BatchInbox = "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f"
	BatchSender = "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f"
	ConcurrentFetchers = 8
	
	[RpcClient.Arbitrum]
	RPCURL = "http://localhost:8545"
	L1RPCURL = "http://localhost:8545"
	BeaconURL = "http://localhost:8545"
	BatchInbox = "0x1c479675ad559DC151F6Ec7ed3FbF8ceE79582B6"
	ConcurrentFetchers = 8

	[RpcClient.Mantle]
	RPCURL = "http://localhost:8545"
	L1RPCURL = "http://localhost:8545"
	BatchStorageAddr = "0xbB9dDB1020F82F93e45DA0e2CFbd27756DA36956"

	[RpcClient.Mock]
	RPCURL = "http://localhost:8545"

[Consensus]
ProposerBLSKeystorePath = "./testutil/vector/config/bls_0.json"
ProposerBLSKeystorePassword = "password_localtest"
ProposerBLSKeystorePasswordPath = ""
RoundInterval = "500ms"
RoundLimit = "30s"
BLSCurve = "BN254"

[Telemetry]
MetricsEnabled = true
MetricsServerPort = "8080"
ServiceName = "lagrange-node"
PrometheusRetentionTime = "60s"
`
