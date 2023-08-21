package optimism

import (
        "os"
	"testing"
	"log"
	"github.com/stretchr/testify/require"
	"math/big"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
)

var eth = os.Getenv("EthereumURL")
var opt = os.Getenv("RPCEndpoint")
var addr = "0xe6dfba0953616bacab0c9a8ecb3a9bba77fc15c0"

func TestL2Output(t *testing.T) {
    ethClient, err := rpc.Dial(eth)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    output,err := getL2OutputAfter(ethClient, common.HexToAddress(addr), big.NewInt(11991348))
    _ = output
    require.NoError(t, err)
 }

func TestL2OutputProof(t *testing.T) {
    cfg := ProofConfig{
        EthEndpoint: eth,
        OptEndpoint: opt,
        L2OutputOracleAddr: "0xe6dfba0953616bacab0c9a8ecb3a9bba77fc15c0" }
    proof,err := GetProof(cfg, 11991348)
    _ = proof
    require.NoError(t, err)
}