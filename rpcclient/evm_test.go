package rpcclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoints(t *testing.T) {
	c, err := NewEvmClient("http://localhost:8545")
	require.NoError(t, err)
	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	hash, err := c.GetBlockHashByNumber(1)
	require.NoError(t, err)
	require.Equal(t, len(hash), 66)
}
