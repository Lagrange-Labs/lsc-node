package rpcclient

import "fmt"

var (
	// ErrBlockNotFound is returned when the block is not found.
	ErrBlockNotFound = fmt.Errorf("block not found")
)

type RpcClient interface {
	GetBlockHashByNumber(blockNumber uint64) (string, error)
	GetChainID() (uint32, error)
}
