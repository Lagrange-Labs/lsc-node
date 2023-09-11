package arbitrum

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var ethSecret = os.Getenv("ETH_RPC")
var optSecret = os.Getenv("ARB_RPC")

func TestL2Hash(t *testing.T) {
	os.Setenv("EthereumURL", ethSecret)
	os.Setenv("RPCEndpoint", optSecret)

	cfg := ProofConfig{
		EthEndpoint: os.Getenv("EthereumURL"),
		ArbEndpoint: os.Getenv("RPCEndpoint"),
		OutboxAddr:  "0x45Af9Ed1D03703e480CE7d328fB684bb67DA5049", //proxy
	}

	hash, err := GetL2Hash(cfg, 31000000) //problem case: 29477028; working case: 32872241
	require.NoError(t, err)
	require.Equal(t, hash, "0x024ff61e1e2d0b8bba438c7d6433c6bf79d773922af6d5cb822c153e2c16a851")
}
