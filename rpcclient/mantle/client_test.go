package mantle

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const GOERLI_BATCHSTORSGE_ADDR = "0xe5d639b1283352f32477a95b5d4109bcf9d4acf3"
const LOCAL_BATCHSTORSGE_ADDR = "0x2f947E51B9A7cF1d6651D0a568261673233ba42b"

func TestEndpoints(t *testing.T) {
	ethURL := os.Getenv("ETH_RPC")
	c, err := NewClient("http://127.0.0.1:8545", ethURL, GOERLI_BATCHSTORSGE_ADDR)
	require.NoError(t, err)
	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	hash, err := c.GetBlockHashByNumber(1)
	require.NoError(t, err)
	require.Equal(t, len(hash), 66)

	num, err := c.GetFinalizedBlockNumber()
	require.NoError(t, err)
	require.Greater(t, num, uint64(0))
}

func TestFinalizedL2BlockNumberMock(t *testing.T) {
	c, err := NewClient("http://127.0.0.1:8545", "http://127.0.0.1:8545", LOCAL_BATCHSTORSGE_ADDR)
	require.NoError(t, err)

	// pre-merge chain does not support this
	num, err := c.GetFinalizedBlockNumber()
	require.NoError(t, err)
	require.Equal(t, num, uint64(math.MaxUint64))
}
