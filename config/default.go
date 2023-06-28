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
CommitteeSCAddress = "0x928BD9d79f8f15af7E955b8cE174996c231D17C1"
BLSPrivateKey = "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a"
ECDSAPrivateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
PullInterval = "2s"

[Store]
BackendType = "mongodb"
DBPath = "mongodb://127.0.0.1:27017"

[Sequencer]
Chain = "arbitrum"
RPCURL = "http://127.0.0.1:8545"
FromBlockNumber = 1

[Governance]
EthereumURL = "http://127.0.0.1:8545"
PrivateKey = "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c"
StakingSCAddress = "0x880359D8789590cA07bBAb78c0AcCfD3b0a77E49"
CommitteeSCAddress = "0x928BD9d79f8f15af7E955b8cE174996c231D17C1"
StakingCheckInterval = "2s"
EvidenceUploadInterval = "2s"

[Consensus]
OperatorAddress = "0x6E654b122377EA7f592bf3FD5bcdE9e8c1B1cEb9"
ProposerPrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
RoundInterval = "5s"
RoundLimit = "30s"
`
