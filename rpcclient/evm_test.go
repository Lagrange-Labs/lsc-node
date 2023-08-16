package rpcclient

import (
        "os"
	"testing"

	"github.com/stretchr/testify/require"
	"math/big"
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

func TestBlockCollector(t *testing.T) {
	c, err := NewEvmClient(os.Getenv("EthereumURL"))
	if err != nil { panic(err) }
	blocks,err := c.GetRawBlockHeaders(big.NewInt(9500000),big.NewInt(9500010))
	if err != nil { panic(err) }
	require.Equal(t, len(blocks), 11)
}