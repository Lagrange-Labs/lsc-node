package optimism

import (
	"testing"
)

func TestBatchFetcher(t *testing.T) {
	// l1RpcURL := "https://eth-goerli.g.alchemy.com/v2/7dXaVUU3a3Upb47u4o7a9USUbxWD6GwJ"
	// // os.Getenv("ETH_RPC")
	// l2RpcURL := "https://opt-goerli.g.alchemy.com/v2/eDvQvfpLttCU0yQyhXj1DrC9fhqnV5kr"
	// // os.Getenv("OPT_RPC")
	// if len(l1RpcURL) == 0 {
	// 	t.Skip("ETH_RPC not set")
	// }

	// cfg := Config{
	// 	RPCURL:           l2RpcURL,
	// 	L1RPCURL:         l1RpcURL,
	// 	BeginBlockNumber: 10212172,
	// 	BatchInbox:       "0xff00000000000000000000000000000000000420",
	// 	BatchSender:      "0x7431310e026B69BFC676C0013E12A1A11411EEc9",
	// }

	// f, err := NewFetcher(&cfg)
	// require.NoError(t, err)
	// require.NoError(t, f.Fetch())
}
