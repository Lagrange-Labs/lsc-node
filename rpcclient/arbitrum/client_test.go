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

	cfg := &Config{
		RPCURL: arbURL,
	}
	c, err := NewClient(cfg)
	require.NoError(t, err)

	_, err = c.GetCurrentBlockNumber()
	require.NoError(t, err)
}
