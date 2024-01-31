package mantle

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const GOERLI_BATCHSTORSGE_ADDR = "0xe5d639b1283352f32477a95b5d4109bcf9d4acf3"
const LOCAL_BATCHSTORSGE_ADDR = "0xbB9dDB1020F82F93e45DA0e2CFbd27756DA36956"

func TestEndpoints(t *testing.T) {
	ethURL := os.Getenv("ETH_RPC")
	if len(ethURL) == 0 {
		t.Skip()
	}
	c, err := NewClient("http://localhost:8545", ethURL, GOERLI_BATCHSTORSGE_ADDR)
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

func TestFinalizedL2BlockNumberMock(t *testing.T) {
	c, err := NewClient("http://localhost:8545", "http://localhost:8545", LOCAL_BATCHSTORSGE_ADDR)
	require.NoError(t, err)

	// pre-merge chain does not support this
	_, err = c.GetL2FinalizedBlockNumber()
	require.NoError(t, err)
}
