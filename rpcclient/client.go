package rpcclient

import "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"

// NewClient creates a new rpc client.
func NewClient(chain, rpcURL string) (types.RpcClient, error) {
	switch chain {
	// case "arbitrum":
	// 	return NewEvmClient(rpcURL)
	// case "optimism":
	// 	return NewEvmClient(rpcURL)
	default:
		return nil, nil
	}
}
