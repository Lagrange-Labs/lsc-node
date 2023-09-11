package evmclient

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoints(t *testing.T) {
	c, err := NewClient("http://localhost:8545")
	require.NoError(t, err)
	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	hash, err := c.GetBlockHashByNumber(1)
	require.NoError(t, err)
	require.Equal(t, len(hash), 66)

	// pre-merge chain does not support this
	num, err := c.GetL2FinalizedBlockNumber()
	require.Equal(t, num, uint64(math.MaxUint64))
	require.NoError(t, err)
}

func TestFinalizedL2BlockNumber(t *testing.T) {
	optURL := os.Getenv("OPT_RPC")
	if optURL == "" {
		t.Skip("OPT_RPC not set")
	}
	c, err := NewClient(optURL)
	require.NoError(t, err)

	cNum, err := c.GetCurrentBlockNumber()
	require.NoError(t, err)

	num, err := c.GetL2FinalizedBlockNumber()
	require.NoError(t, err)
	require.True(t, num < cNum)

	arbURL := os.Getenv("ARB_RPC")
	if arbURL == "" {
		t.Skip("ARB_RPC not set")
	}
	c, err = NewClient(arbURL)
	require.NoError(t, err)

	cNum, err = c.GetCurrentBlockNumber()
	require.NoError(t, err)

	num, err = c.GetL2FinalizedBlockNumber()
	require.NoError(t, err)
	require.True(t, num < cNum)
}
