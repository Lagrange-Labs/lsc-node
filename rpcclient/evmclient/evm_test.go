package evmclient

import (
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

	l2Hash, err := c.GetBlockHashByNumber(1) //nolint:staticcheck
	require.NoError(t, err)
	require.Equal(t, len(l2Hash), 32)

	// pre-merge chain does not support this
	_, err = c.GetFinalizedBlockNumber()
	require.NoError(t, err)

	header, err := c.GetRawHeaderByNumber(1)
	require.NoError(t, err)
	require.NotNil(t, header)
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

	num, err := c.GetFinalizedBlockNumber()
	require.NoError(t, err)
	require.Greater(t, num, uint64(0))
	require.True(t, num < cNum)

	arbURL := os.Getenv("ARB_RPC")
	if arbURL == "" {
		t.Skip("ARB_RPC not set")
	}
	c, err = NewClient(arbURL)
	require.NoError(t, err)

	cNum, err = c.GetCurrentBlockNumber()
	require.NoError(t, err)

	num, err = c.GetFinalizedBlockNumber()
	require.NoError(t, err)
	require.Greater(t, num, uint64(0))
	require.True(t, num < cNum)
}

func TestBlocksByRange(t *testing.T) {
	c, err := NewClient("http://localhost:8545")
	require.NoError(t, err)

	start := uint64(1)
	end := uint64(10)
	hashes, err := c.GetBlockHashesByRange(start, end)
	require.NoError(t, err)
	require.Len(t, hashes, int(end-start))
}
