package rpcclient

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/arbitrum"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mantle"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mock"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/optimism"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

// NewClient creates a new rpc client.
func NewClient(chain string, cfg *Config) (types.RpcClient, error) {
	switch chain {
	case "mantle":
		return mantle.NewClient(cfg.Mantle)
	case "arbitrum":
		return arbitrum.NewClient(cfg.Arbitrum)
	case "optimism":
		return optimism.NewClient(cfg.Optimism)
	case "mock":
		return mock.NewClient(cfg.Mock)
	default:
		return nil, nil
	}
}
