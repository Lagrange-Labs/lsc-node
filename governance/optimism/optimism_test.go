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

var ethSecret = os.Getenv("ETH_RPC")
var optSecret = os.Getenv("OPT_RPC")

var addr = "0xe6dfba0953616bacab0c9a8ecb3a9bba77fc15c0"

func TestL2Output(t *testing.T) {
    os.Setenv("EthereumURL",ethSecret)
    os.Setenv("RPCEndpoint",optSecret)

    eth := os.Getenv("EthereumURL")

    ethClient, err := rpc.Dial(eth)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    output,err := getL2OutputAfter(ethClient, common.HexToAddress(addr), big.NewInt(11991348))
    _ = output
    if err != nil {
        log.Fatalf("%v",err)
    }
    require.NoError(t, err)
 }

func TestL2OutputProof(t *testing.T) {
    os.Setenv("EthereumURL",ethSecret)
    os.Setenv("RPCEndpoint",optSecret)

    eth := os.Getenv("EthereumURL")
    opt := os.Getenv("RPCEndpoint")
    
    cfg := ProofConfig{
        EthEndpoint: eth,
        OptEndpoint: opt,
        L2OutputOracleAddr: "0xe6dfba0953616bacab0c9a8ecb3a9bba77fc15c0" }
    proof,err := GetProof(cfg, 11991348)
    _ = proof
    if err != nil {
        log.Fatalf("%v",err)
    }
    require.NoError(t, err)
}