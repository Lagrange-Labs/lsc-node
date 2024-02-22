package mantle

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpointsRemote(t *testing.T) {
	ethURL := os.Getenv("ETH_RPC")
	if len(ethURL) == 0 {
		t.Skip()
	}
	rpcURL := os.Getenv("MANTLE_RPC_URL")
	if len(rpcURL) == 0 {
		t.Skip()
	}
	l2URL := strings.Replace(rpcURL, "9545", "8545", 1)
	c, err := NewClient(l2URL, ethURL, rpcURL)
	require.NoError(t, err)

	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	hash, err := c.GetBlockHashByNumber(1)
	require.NoError(t, err)
	require.Equal(t, len(hash), 66)

	num, err := c.GetL2FinalizedBlockNumber()
	require.NoError(t, err)
	require.Greater(t, num, uint64(0))
}
