package mantle

import (
    "os"
    "fmt"
    "strings"
    //"context"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    ethclient "github.com/ethereum/go-ethereum/ethclient"
)

const (
    rollupAddress   = "0xb7d41213822cc662a6de97f53a0c58b33c179553"
    assertionMapAddress = "0x14b3485C6d07Ca7b16C73d595e9B0a07362Ae90b"
)

func GetClient() *ethclient.Client {
    eth,err := ethclient.Dial(os.Getenv("ETH_RPC"))
    if err != nil { panic(err) }
    return eth
}

func GetAbi(abiStr string) abi.ABI {
    abiParse, err := abi.JSON(strings.NewReader(abiStr))    
    if err != nil { panic(err) }
    return abiParse
}

func GetCheckpointAndRootByBlockNumber(num int) string {
    // Initialize vars
    eth := GetClient()

    rollupAddr := common.HexToAddress(rollupAddress)
    rollup,err := NewRollup(rollupAddr,eth)
    if err != nil { panic(err) }

    assertAddr := common.HexToAddress(assertionMapAddress)
    assertionMap,err := NewAssertionMap(assertAddr,eth)
    if err != nil { panic(err) }
    
    var root [32]byte
    var l2BlockNum *big.Int
    var assertID *big.Int

    // Get latest assertion ID
    assertID,err = rollup.LastConfirmedAssertionID(&bind.CallOpts{})
    if err != nil { panic(err) }

    // Get latest assertion corresponding to ID
    assertion, err := assertionMap.Assertions(&bind.CallOpts{},assertID)
    if err != nil { panic(err) }
    
    for {
        fmt.Println(assertID)
        // Retrieve relevant values from assertion
        l2BlockNum = assertion.InboxSize
        root = common.HexToHash(hexutil.Encode(assertion.StateHash[:]))

    // Generate ABI for evidence extraData
    abiParse := GetAbi(`[
    {
        "constant": false,
        "inputs": [
            {
                "name": "assertionID",
                "type": "uint256"
            },
            {
                "name": "root",
                "type": "bytes32"
            },
            {
                "name": "l2blocknum",
                "type": "uint256"
            }
        ],
        "name": "rootAndBlocknum",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`)


        // Decrement assertion index
        assertID.Sub(assertID,big.NewInt(1))
        
        // Assertion lookahead
        newAssertion, err := assertionMap.Assertions(&bind.CallOpts{},assertID)
        if err != nil { panic(err) }
        
        // If previous assertion is less than block num, break with relevant data
        if newAssertion.InboxSize.Cmp(big.NewInt(int64(num))) == -1 {
            break;
        }
        
        // Otherwise, rinse and repeat
        assertion = newAssertion
    }
    
    packedBytes, err := abiParse.Pack("rootAndBlocknum", assertionID, root, l2BlockNum)
    if err != nil { panic(err) }
    
    fmt.Println(root, l2BlockNum)
    return hexutil.Encode(packedBytes)
}