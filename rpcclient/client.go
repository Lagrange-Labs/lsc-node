package rpcclient

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mantle"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

// NewClient creates a new rpc client.
func NewClient(chain, rpcURL, ethURL, batchStorage string) (types.RpcClient, error) {
	switch chain {
	case "mantle":
		return mantle.NewClient(rpcURL, ethURL, batchStorage)
	case "arbitrum":
		return evmclient.NewClient(rpcURL)
	case "optimism":
		return evmclient.NewClient(rpcURL)
	default:
		return nil, nil
	}
}
