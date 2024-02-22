package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"

[Client]
GrpcURL = "127.0.0.1:9090"
Chain = "arbitrum"
RPCEndpoint = "http://localhost:8545"
EthereumURL = "http://localhost:8545"
CommitteeSCAddress = "0xF2740f6A6333c7B405aD7EfC68c74adAd83cC30D"
BLSPrivateKey = "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a"
ECDSAPrivateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
PullInterval = "100ms"
BLSCurve = "BN254"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://127.0.0.1:27017"

[Sequencer]
Chain = "arbitrum"
RPCURL = "http://localhost:8545"
EthURL = "http://localhost:8545"
NewRPCURL = "http://localhost:8545"
FromBlockNumber = 5

[Governance]
EthereumURL = "http://localhost:8545"
PrivateKey = "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c"
CommitteeSCAddress = "0xF2740f6A6333c7B405aD7EfC68c74adAd83cC30D"
StakingCheckInterval = "2s"
EvidenceUploadInterval = "3s"

[Consensus]
OperatorAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
ProposerPrivateKey = "0x8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b"
BatchSize = 20
RoundInterval = "500ms"
RoundLimit = "30s"
BLSCurve = "BN254"
`
