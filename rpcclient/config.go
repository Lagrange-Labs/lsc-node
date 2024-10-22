package rpcclient

import (
	"github.com/Lagrange-Labs/lsc-node/rpcclient/arbitrum"
	"github.com/Lagrange-Labs/lsc-node/rpcclient/mantle"
	"github.com/Lagrange-Labs/lsc-node/rpcclient/mock"
	"github.com/Lagrange-Labs/lsc-node/rpcclient/optimism"
)

// Config is a config for rpc client.
type Config struct {
	Optimism *optimism.Config `yaml:"Optimism"`
	Mantle   *mantle.Config   `yaml:"Mantle"`
	Arbitrum *arbitrum.Config `yaml:"Arbitrum"`
	Mock     *mock.Config     `yaml:"Mock"`
}
