package rpcclient

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

// NewClient creates a new rpc client.
func NewClient(chain, rpcURL string) (types.RpcClient, error) {
	switch chain {
	case "arbitrum":
		return evmclient.NewClient(rpcURL)
	case "optimism":
		return evmclient.NewClient(rpcURL)
	default:
		return nil, nil
	}
}
