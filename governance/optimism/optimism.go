package optimism

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"os/exec"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	rlp "github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

// OutputRootProof is derived from from Optimism L1 Contract Types
type OutputRootProof struct {
	Version                  *big.Int
	StateRoot                common.Hash
	MessagePasserStorageRoot common.Hash
	LatestBlockhash          common.Hash
}

// OutputProposal is derived from from Optimism L1 Contract Types
type OutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// ProofConfig is the endpoint/address config for reconstructing and verifying outputs
type ProofConfig struct {
	EthEndpoint        string
	OptEndpoint        string
	L2OutputOracleAddr string
}

func getL2OutputAfter(rpc *rpc.Client, addr common.Address, blockNum *big.Int) (OutputProposal, error) {
	// Load and Parse ABI
	abiPath := "../../scinterface/bin/goerli/L2OutputOracle.json"
	abiJSON, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return OutputProposal{}, err
	}
	l2ooAbi, err := abi.JSON(strings.NewReader(string(abiJSON)))
	if err != nil {
		return OutputProposal{}, err
	}

	// Make RPC Request for L2 Output Proposal
	f := "getL2OutputAfter"
	args := []interface{}{blockNum}
	data, err := l2ooAbi.Pack(f, args...)
	if err != nil {
		return OutputProposal{}, err
	}

	type request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	call := request{addr.Hex(), hexutil.Encode(data)}

	var res string
	err = rpc.Call(&res, "eth_call", call, "latest")
	if err != nil {
		return OutputProposal{}, err
	}

	// Process RPC result and convert to OutputProposal struct

	resBytes, err := hexutil.Decode(res)
	if err != nil {
		return OutputProposal{}, err
	}

	values, err := l2ooAbi.Unpack("getL2OutputAfter", resBytes)
	if err != nil {
		return OutputProposal{}, err
	}

	op := values[0]
	outputProposal := op.(struct {
		OutputRoot    [32]uint8 "json:\"outputRoot\""
		Timestamp     *big.Int  "json:\"timestamp\""
		L2BlockNumber *big.Int  "json:\"l2BlockNumber\""
	})

	return OutputProposal{
		common.Hash(outputProposal.OutputRoot),
		outputProposal.Timestamp,
		outputProposal.L2BlockNumber}, nil
}

// Hex - converts output root proof to hex string
func (orp *OutputRootProof) Hex() (string, error) {
	proofABI, err := abi.JSON(strings.NewReader(`[{"type": "function", "name": "bytes32[4]", "inputs": [{"name": "a", "type": "uint256"},{"name": "b", "type": "bytes32"},{"name": "c", "type": "bytes32"},{"name": "d", "type": "bytes32"}]}]`))
	if err != nil {
		return "", err
	}
	encoded, err := proofABI.Pack(
		"bytes32[4]",
		orp.Version,
		orp.StateRoot,
		orp.MessagePasserStorageRoot,
		orp.LatestBlockhash,
	)
	if err != nil {
		return "", err
	}
	encodedCleaned := encoded[4:] // strip signature prefix
	return hexutil.Encode(encodedCleaned), nil
}

// GetProof - Reconstructs nearet OutputRootProof immediately following provided blockNumber
func GetProof(cfg ProofConfig, blockNumber int) (OutputRootProof, error) {
	blockNum := big.NewInt(int64(blockNumber))
	_ = blockNum

	// Initialize RPC and ETH Clients
	ethClient, err := rpc.Dial(cfg.EthEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	optRPC, err := rpc.Dial(cfg.OptEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Optimism client: %v", err)
	}

	optClient := ethclient.NewClient(optRPC)

	// Get L2 Output Proposal for block, process
	output, err := getL2OutputAfter(ethClient, common.HexToAddress(cfg.L2OutputOracleAddr), blockNum)
	if err != nil {
		return OutputRootProof{}, err
	}

	outputRoot := output.OutputRoot
	l2BlockNumber := output.L2BlockNumber
	outputRootStr := hexutil.Encode(outputRoot[:])

	// Retrieve L2 block, continue accumulating output proof components
	block, err := optClient.HeaderByNumber(context.Background(), l2BlockNumber)
	if err != nil {
		log.Fatal(err)
	}

	rlpBytes, err := rlp.EncodeToBytes(block)
	if err != nil {
		return OutputRootProof{}, err
	}
	hash := hexutil.Encode(crypto.Keccak256(rlpBytes))

	stateRoot := block.Root
	version := 0

	// Get Storage Root of MessagePasser contract at L2 block
	messagePasserAddr := "0x4200000000000000000000000000000000000016"

	req := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getProof","params":["%s",[],"%s"],"id":1}`, messagePasserAddr, fmt.Sprintf("0x%X", l2BlockNumber.Int64()))

	cmd := exec.Command("curl", "-X", "POST", "--data", req, cfg.OptEndpoint)

	scres, err := cmd.Output()
	if err != nil {
		return OutputRootProof{}, err
	}

	// Parse storage proof JSON

	jsonStr := string(scres)
	var result map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return OutputRootProof{}, err
	}

	if curlResult, ok := result["result"].(map[string]interface{}); ok {
		storageRoot := curlResult["storageHash"]

		// Reconstruct output proof, compute hash, verify against proposal, return

		proofABI, err := abi.JSON(strings.NewReader(`[{"type": "function", "name": "bytes32[4]", "inputs": [{"name": "a", "type": "uint256"},{"name": "b", "type": "bytes32"},{"name": "c", "type": "bytes32"},{"name": "d", "type": "bytes32"}]}]`))
		if err != nil {
			return OutputRootProof{}, err
		}
		encoded, err := proofABI.Pack("bytes32[4]", big.NewInt(int64(version)), stateRoot, common.HexToHash(storageRoot.(string)), common.HexToHash(hash))
		encodedCleaned := encoded[4:] // strip signature prefix
		if err != nil {
			return OutputRootProof{}, err
		}
		reProof := crypto.Keccak256(encodedCleaned)

		if hexutil.Encode(reProof) != outputRootStr {
			return OutputRootProof{}, fmt.Errorf("Output roots do not match: %v %v", hexutil.Encode(reProof), outputRootStr)
		}
		//fmt.Println("Reconstructed Output Root:", hexutil.Encode(reProof))

		proof := OutputRootProof{
			big.NewInt(int64(version)),
			stateRoot,
			common.HexToHash(storageRoot.(string)),
			common.HexToHash(hash)}
		return proof, nil
	}
	return OutputRootProof{}, fmt.Errorf("Failed to parse storage proof: %s", jsonStr)
}
