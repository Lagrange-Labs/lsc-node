package types

import "fmt"

var (
	// ErrBlockNotFound is returned when the block is not found.
	ErrBlockNotFound = fmt.Errorf("block not found")
	ErrUnsupportedNetwork = fmt.Errorf("unsupported network")
)

type RpcClient interface {
	GetCurrentBlockNumber() (uint64, error)
	GetBlockHashByNumber(blockNumber uint64) (string, error)
	GetL2FinalizedBlockNumber() (uint64, error)
	GetChainID() (uint32, error)
}
