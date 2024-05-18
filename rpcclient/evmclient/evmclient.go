package evmclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/lru"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CacheSize = 2048

var _ types.EvmClient = (*Client)(nil)

// Client is an EVM client.
type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	rpcURL    string

	cache *lru.Cache[uint64, json.RawMessage] // block number -> raw header
}

// NewClient creates a new EvmClient instance.
func NewClient(rpcURL string) (*Client, error) {
	client, err := rpc.DialContext(utils.GetContext(), rpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcClient: client,
		ethClient: ethclient.NewClient(client),
		rpcURL:    rpcURL,
		cache:     lru.NewCache[uint64, json.RawMessage](CacheSize),
	}, nil
}

// GetEthClient returns the ethclient.Client.
func (c *Client) GetEthClient() *ethclient.Client {
	return c.ethClient
}

// GetCurrentBlockNumber returns the current block number.
func (c *Client) GetCurrentBlockNumber() (uint64, error) {
	header, err := c.ethClient.HeaderByNumber(utils.GetContext(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *Client) GetBlockHashByNumber(blockNumber uint64) (common.Hash, error) {
	rawHeader, err := c.GetRawHeaderByNumber(blockNumber)
	if errors.Is(err, rpc.ErrNoResult) {
		return common.Hash{}, types.ErrBlockNotFound
	}
	if err != nil {
		logger.Errorf("failed to get the raw header error: %v", err)
		return common.Hash{}, fmt.Errorf("failed to get the raw header error: %w", err)
	}
	return getHashFromRawHeader(rawHeader)
}

func getHashFromRawHeader(rawHeader json.RawMessage) (common.Hash, error) {
	result := &struct {
		Hash common.Hash `json:"hash"`
	}{}
	if err := json.Unmarshal(rawHeader, &result); err != nil {
		return common.Hash{}, fmt.Errorf("failed to unmarshal block header error: %w rawHeader: %s", err, rawHeader)
	}

	return result.Hash, nil
}

// GetBlockNumberByHash returns the block number by the given block hash.
func (c *Client) GetBlockNumberByHash(blockHash common.Hash) (uint64, error) {
	header, err := c.ethClient.HeaderByHash(utils.GetContext(), blockHash)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// GetBlockNumberByTxHash returns the block number by the given transaction hash.
func (c *Client) GetBlockNumberByTxHash(txHash common.Hash) (uint64, error) {
	receipt, err := c.ethClient.TransactionReceipt(utils.GetContext(), txHash)
	if err != nil {
		return 0, err
	}
	return receipt.BlockNumber.Uint64(), nil
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() (uint32, error) {
	chainID, err := c.ethClient.ChainID(utils.GetContext())
	if err != nil {
		return 0, err
	}
	return uint32(chainID.Int64()), err
}

// GetBlockHashesByRange returns the block hashes by the given block number range.
// The range is [start, end).
func (c *Client) GetBlockHashesByRange(start, end uint64) ([]common.Hash, error) {
	batchElems := make([]rpc.BatchElem, 0, int(end-start))
	for i := start; i < end; i++ {
		batchElems = append(batchElems, rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{hexutil.EncodeBig(big.NewInt(int64(i))), false},
			Result: &json.RawMessage{},
		})
	}
	err := c.rpcClient.BatchCallContext(utils.GetContext(), batchElems)
	if err != nil {
		return nil, err
	}
	hashes := make([]common.Hash, 0, len(batchElems))
	for i, batch := range batchElems {
		if batch.Error != nil {
			return nil, batch.Error
		}
		rawHeader := batch.Result.(*json.RawMessage)
		c.cache.Add(start+uint64(i), *rawHeader)
		hash, err := getHashFromRawHeader(*rawHeader)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, hash)
	}
	return hashes, nil
}

// GetFinalizedBlockNumber returns the finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	var header *ethtypes.Header
	if err := c.rpcClient.CallContext(utils.GetContext(), &header, "eth_getBlockByNumber", "finalized", false); err != nil {
		if strings.Contains(err.Error(), "'finalized' tag not supported on pre-merge network") {
			err := c.rpcClient.CallContext(utils.GetContext(), &header, "eth_getBlockByNumber", "latest", false)
			if err != nil {
				logger.Errorf("failed to get latest block number error: %v", err)
				return 0, err
			}
			return header.Number.Uint64(), nil
		}
		logger.Errorf("failed to get finalized block number error: %v", err)
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// GetRawHeaderByNumber returns the raw message of block header by the given block number.
func (c *Client) GetRawHeaderByNumber(blockNumber uint64) (json.RawMessage, error) {
	if raw, ok := c.cache.Get(blockNumber); ok {
		return raw, nil
	}

	var rawHeader json.RawMessage
	if err := c.rpcClient.CallContext(utils.GetContext(), &rawHeader, "eth_getBlockByNumber",
		hexutil.EncodeBig(big.NewInt(int64(blockNumber))), false); err != nil {
		if errors.Is(err, rpc.ErrNoResult) {
			return nil, types.ErrNoResult
		}
		return nil, err
	}
	if len(rawHeader) < 5 { // to detect empty response and "null" response
		return nil, types.ErrNoResult
	}

	c.cache.Add(blockNumber, rawHeader)

	return rawHeader, nil
}

// GetBlockByNumber returns the block header by the given block number.
func (c *Client) GetBlockByNumber(blockNumber uint64) (*ethtypes.Block, error) {
	return c.ethClient.BlockByNumber(utils.GetContext(), big.NewInt(int64(blockNumber)))
}
