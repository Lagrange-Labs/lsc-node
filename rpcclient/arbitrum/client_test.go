package arbitrum

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestL1BlockNumber(t *testing.T) {
	arbURL := os.Getenv("ARB_RPC")
	if arbURL == "" {
		t.Skip("ARB_RPC not set")
	}

	cfg := &Config{
		RPCURL: arbURL,
	}
	c, err := NewClient(cfg)
	require.NoError(t, err)

	cNum, err := c.GetCurrentBlockNumber()
	require.NoError(t, err)

	header, err := c.GetBlockHeaderByNumber(cNum, common.Hash{})
	require.NoError(t, err)
	require.Greater(t, header.L1BlockNumber, uint64(0))
}
