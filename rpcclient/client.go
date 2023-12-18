package rpcclient

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/arbitrum"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mantle"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/optimism"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

// NewClient creates a new rpc client.
func NewClient(chain, rpcURL, ethURL, batchStorageAddr string) (types.RpcClient, error) {
	switch chain {
	case "mantle":
		return mantle.NewClient(rpcURL, ethURL, batchStorageAddr)
	case "arbitrum":
		return arbitrum.NewClient(rpcURL, ethURL, batchStorageAddr)
	case "optimism":
		return optimism.NewClient(&optimism.Config{})
	default:
		return nil, nil
	}
}
