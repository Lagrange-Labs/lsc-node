package evmclient

import (
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
)

type Client struct {
	ethClient *ethclient.Client
	rpcURL    string
}

var _ types.RpcClient = (*Client)(nil)

// NewClient creates a new EvmClient instance.
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

// GetChainID returns the chain ID.
func (c *Client) GetChainID() (uint32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	return uint32(chainID.Int64()), err
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
