package rpcclient

import (
	"fmt"
	"strings"

	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/arbitrum"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mantle"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/mock"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/optimism"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

// NewClient creates a new rpc client.
func NewClient(chain string, cfg *Config, isLight bool) (types.RpcClient, error) {
	logger.Infof("creating the rpc client for chain: %s", chain)

	switch strings.ToLower(chain) {
	case "mantle":
		return mantle.NewClient(cfg.Mantle, isLight)
	case "arbitrum":
		return arbitrum.NewClient(cfg.Arbitrum, isLight)
	case "base":
		return optimism.NewClient(cfg.Optimism, isLight)
	case "optimism":
		return optimism.NewClient(cfg.Optimism, isLight)
	case "polymer":
		return optimism.NewClient(cfg.Optimism, isLight)
	case "mock":
		return mock.NewClient(cfg.Mock, isLight)
	default:
		return nil, fmt.Errorf("unsupported chain: %s", chain)
	}
}
