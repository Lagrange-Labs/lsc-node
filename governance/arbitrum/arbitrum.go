package arbitrum

import (
    "context"
    "math/big"
    "strings"
    "io/ioutil"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

type ProofConfig struct {
    EthEndpoint		string
    ArbEndpoint		string
    OutboxAddr		string
}

func GetL2Hash(cfg ProofConfig, blockNumber int) (string, error) {
    blockNum := big.NewInt(int64(blockNumber))

    // Initialize RPC clients
    eth_rpc, err := rpc.Dial(cfg.EthEndpoint)
    if err != nil { return "", err }
    arb_client, err := ethclient.Dial(cfg.ArbEndpoint)
    if err != nil { return "", err }

    // Get SendRoot from L2 block header
    header, err := arb_client.HeaderByNumber(context.Background(), blockNum)
    if err != nil { return "", err }
    extraData := header.Extra

    // Load and parse ABI
    abiPath := "../../scinterface/bin/goerli/Outbox.json"
    abiJSON, err := ioutil.ReadFile(abiPath)
    if err != nil { return "",err }
    outbox_abi, err := abi.JSON(strings.NewReader(string(abiJSON)))
    if err != nil { return "",err }

    // Prepare and make RPC request to outbox for checkpoint block hash
    f := "roots"
    data, err := outbox_abi.Pack(f, common.BytesToHash(extraData))
    if err != nil { return "",err }
    
    type request struct {
        To   string `json:"to"`
        Data string `json:"data"`
    }
    addr := cfg.OutboxAddr
    call := request{addr,hexutil.Encode(data)}

    var res string
    err = eth_rpc.Call(&res, "eth_call", call, "latest")
    if err != nil { return "",err }
    
    resBytes,err := hexutil.Decode(res)
    if err != nil { return "",err }

    values, err := outbox_abi.Unpack("roots", resBytes)
    if err != nil { return "",err }
    
    op := values[0]
    checkpointBlockHash := op.([32]uint8)
    return hexutil.Encode(checkpointBlockHash[:]),nil
}