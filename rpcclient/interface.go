package rpcclient

type RpcClient interface {
	GetBlockHashByNumber(blockNumber int64) (string, error)
	GetChainID() (int32, error)
}
