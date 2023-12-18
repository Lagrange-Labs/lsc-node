package rpcclient

import "github.com/Lagrange-Labs/lagrange-node/rpcclient/optimism"

// RpcClientConfig is a config for rpc client.
type RpcClientConfig struct {
	Optimism *optimism.Config `yaml:"optimism"`
}
