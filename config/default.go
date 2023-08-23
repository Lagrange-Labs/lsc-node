package config

// DefaultValues is the default configuration
const DefaultValues = `
[Server]
GRPCPort = "9090"

[Client]
GrpcURL = "127.0.0.1:9090"
Chain = "arbitrum"
RPCEndpoint = "http://127.0.0.1:8545"
EthereumURL = "http://127.0.0.1:8545"
CommitteeSCAddress = "0x923d8ADAAa6e52c485293cD48EE56F7BFAD85cd4"
BLSPrivateKey = "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a"
ECDSAPrivateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
PullInterval = "100ms"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://127.0.0.1:27017"

[Sequencer]
Chain = "arbitrum"
RPCURL = "http://127.0.0.1:8545"
FromBlockNumber = 5

[Governance]
EthereumURL = "http://127.0.0.1:8545"
PrivateKey = "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c"
StakingSCAddress = "0xF824C350EA9501234a731B01B8EC6E660e069c7F"
CommitteeSCAddress = "0x923d8ADAAa6e52c485293cD48EE56F7BFAD85cd4"
StakingCheckInterval = "2s"
EvidenceUploadInterval = "3s"
L2OutputOracle = "0xe6dfba0953616bacab0c9a8ecb3a9bba77fc15c0"
Outbox = "0x45Af9Ed1D03703e480CE7d328fB684bb67DA5049"

[Consensus]
OperatorAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
ProposerPrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
BatchSize = 20
RoundInterval = "500ms"
RoundLimit = "30s"
`
