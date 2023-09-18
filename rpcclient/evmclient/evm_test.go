package evmclient

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
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

func TestBlockCollector(t *testing.T) {
	c, err := NewClient(os.Getenv("ETH_RPC"))
	if err != nil { panic(err) }
	blocks,err := c.GetRawBlockHeaders(big.NewInt(9500000),big.NewInt(9500010))
	if err != nil { panic(err) }
	require.Equal(t, len(blocks), 11)
}

func TestRawBlockHeaders(t *testing.T) {
	c, err := NewClient(os.Getenv("ETH_RPC"))
	if err != nil { panic(err) }
	block,err := c.getRawBlockHeader(1)
	if err != nil { panic(err) }
	blockBytes,err := hexutil.Decode(block)
	if err != nil { panic(err) }
	// Eth Goerli Assumed
	hash := crypto.Keccak256Hash(blockBytes)
	require.Equal(t, "0x8f5bab218b6bb34476f51ca588e9f4553a3a7ce5e13a66c660a5283e97e9a85a", hash.Hex())
}
func TestRawAttestBlockHeaders(t *testing.T) {
        os.Setenv("RPCEndpoint",os.Getenv("OPT_RPC"))
	block,err := GetRawAttestBlockHeader(1)
	if err != nil { panic(err) }
	blockBytes,err := hexutil.Decode(block)
	if err != nil { panic(err) }
	// Opt Goerli Assumed
	hash := crypto.Keccak256Hash(blockBytes)
	require.Equal(t, "0x15d55041e8f7b0d1f303b6d4cefe2d2efc257d67acd9f17307261a8f7d786e0e", hash.Hex())
}
