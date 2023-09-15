package mantle

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoints(t *testing.T) {
	ethURL := os.Getenv("ETH_RPC")
	c, err := NewClient("http://localhost:8545", ethURL, "0xe5d639b1283352f32477a95b5d4109bcf9d4acf3")
	require.NoError(t, err)
	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	hash, err := c.GetBlockHashByNumber(1)
	require.NoError(t, err)
	require.Equal(t, len(hash), 66)

	// pre-merge chain does not support this
	num, err := c.GetL2FinalizedBlockNumber()
	require.NoError(t, err)
	require.Greater(t, num, uint64(0))
}
