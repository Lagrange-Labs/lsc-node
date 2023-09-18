package evmclient

import (
	"os"
	"context"
	"encoding/json"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
        rlp "github.com/ethereum/go-ethereum/rlp"
        common "github.com/ethereum/go-ethereum/common"
        hexutil "github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Lagrange-Labs/lagrange-node/governance/arbitrum"
	"github.com/Lagrange-Labs/lagrange-node/governance/optimism"
)

type Client struct {
	ethClient *ethclient.Client
	rpcURL    string
}

var _ types.RpcClient = (*Client)(nil)

// NewClient creates a new Client instance.
func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		ethClient: client,
		rpcURL:    rpcURL,
	}, nil
}

// GetCurrentBlockNumber returns the current block number.
func (c *Client) GetCurrentBlockNumber() (uint64, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *Client) GetBlockHashByNumber(blockNumber uint64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err == ethereum.NotFound {
		return "", types.ErrBlockNotFound
	}

	return header.Hash().Hex(), err
}

// GetExtraDataByNumber returns the block extradata by the given block number.
func (c *Client) GetExtraDataByNumber(blockNumber uint64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err == ethereum.NotFound {
		return "", types.ErrBlockNotFound
	}
	extraData := header.Extra

	return hexutil.Encode(extraData), err
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() (uint32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	return uint32(chainID.Int64()), err
}

// GetRawAttestBlockHeader returns the raw block header hex string associated with blockNum w/o explicit client
func GetRawAttestBlockHeader(blockNum int) (string, error) {
	optClient,err := NewClient(os.Getenv("RPCEndpoint"))
	if err != nil { return "0x00",nil }
	hex,err := optClient.getRawBlockHeader(blockNum)
	return hex,err
}

// getRawBlockHeader returns the raw block header hex string associated with blockNum
func (c *Client) getRawBlockHeader(blockNum int) (string, error) {
        header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil { return "",err }
	rlpBytes, err := rlp.EncodeToBytes(header)
	if err != nil { return "",err }	
        hex := hexutil.Encode(rlpBytes)
	return hex,nil
}

// GetRawBlockHeaders retrieves RLP-encoded block headers, boundary-inclusive.
func (c *Client) GetRawBlockHeaders(startblock *big.Int, endblock *big.Int) (map[*big.Int]string, error) {
    headers := make(map[*big.Int]string)
    // Iterate block numbers
    for i := (*startblock).Int64(); i <= (*endblock).Int64(); i++ {
        hex,err := c.getRawBlockHeader(int(i))
        if err != nil {
	    return headers,err
	}
        // Collect raw block header
        headers[big.NewInt(i)] = hex
    }
    return headers,nil
}

// GetExtraDataByNetwork returns blockNum header's extradata field as hex string
func GetExtraDataByNetwork(blockNum int) (string, common.Hash, error) {
	network := os.Getenv("Chain")
	if network == "arbitrum" {
	    proofCfg := arbitrum.ProofConfig{
	        EthEndpoint: os.Getenv("EthereumURL"),
		ArbEndpoint: os.Getenv("RPCendpoint"),
		OutboxAddr: os.Getenv("Outbox"),
	    }
	    l2Hash,err := arbitrum.GetL2Hash(proofCfg, blockNum)
	    if err != nil { return "0x00", common.HexToHash("0x00"), err }
	    eth,err := NewClient(os.Getenv("EthereumURL"))
	    if err != nil { return "0x00", common.HexToHash("0x00"), err }
	    extra,err := eth.GetExtraDataByNumber(uint64(blockNum))
	    if err != nil { return "0x00", common.HexToHash("0x00"), err }
	    return extra, common.HexToHash(l2Hash), nil
	} else if network == "optimism" {
	    proofCfg := optimism.ProofConfig{
	        EthEndpoint: os.Getenv("EthereumURL"),
		OptEndpoint: os.Getenv("RPCendpoint"),
		L2OutputOracleAddr: os.Getenv("L2OutputOracle"),
	    }
	    proof, err := optimism.GetProof(proofCfg, blockNum)
	    if err != nil { return "0x00", common.HexToHash("0x00"), err }
	    proofHex, err := proof.Hex()
	    if err != nil { return "0x00", common.HexToHash("0x00"), err }
	    return proofHex, proof.LatestBlockhash, nil
	} else {
	    return "0x00", common.HexToHash("0x00"), types.ErrUnsupportedNetwork
	}
}

// GetL2FinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetL2FinalizedBlockNumber() (uint64, error) {
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"finalized\",false]}")

	req, _ := http.NewRequest("POST", c.rpcURL, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var result = struct {
		Result struct {
			Number string `json:"number"`
		} `json:"result"`
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Errorf("failed to unmarshal json: %v", string(body))
		return 0, err
	}
	if result.Error.Code != 0 {
		// TODO: handle error
		// API does not support the finalized block number
		return math.MaxUint64, nil
	}
	blockNumber, err := strconv.ParseUint(result.Result.Number, 0, 64)
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}
