package arbitrum

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestL1BlockNumber(t *testing.T) {
	arbURL := os.Getenv("ARB_RPC")
	if arbURL == "" {
		t.Skip("ARB_RPC not set")
	}

	c, err := NewClient(arbURL, arbURL, "")
	require.NoError(t, err)

	cNum, err := c.GetCurrentBlockNumber()
	require.NoError(t, err)

	header, err := c.GetBlockHeaderByNumber(cNum)
	require.NoError(t, err)
	require.Greater(t, header.L1BlockNumber, uint64(0))
}
