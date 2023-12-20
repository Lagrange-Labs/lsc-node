package evmclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/lru"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

const CacheSize = 128

// Client is an EVM client.
type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	rpcURL    string

	cache *lru.Cache[uint64, json.RawMessage] // block number -> raw header
}

// NewClient creates a new EvmClient instance.
func NewClient(rpcURL string) (*Client, error) {
	client, err := rpc.DialContext(context.Background(), rpcURL)
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

// GetCurrentBlockNumber returns the current block number.
func (c *Client) GetCurrentBlockNumber() (uint64, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *Client) GetBlockHashByNumber(blockNumber uint64) (common.Hash, error) {
	rawHeader, err := c.GetRawHeaderByNumber(blockNumber)
	if err == rpc.ErrNoResult {
		return common.Hash{}, types.ErrBlockNotFound
	}
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get the raw header error: %w", err)
	}

	var header ethtypes.Header
	if err := json.Unmarshal(rawHeader, &header); err != nil {
		return common.Hash{}, fmt.Errorf("failed to unmarshal block header error: %w rawHeader: %s", err, rawHeader)
	}

	return header.Hash(), err
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() (uint32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	return uint32(chainID.Int64()), err
}

// GetFinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	var header *ethtypes.Header
	if err := c.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", "finalized", false); err != nil {
		if strings.Contains(err.Error(), "'finalized' tag not supported on pre-merge network") {
			return math.MaxUint64, nil
		}
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
	if err := c.rpcClient.CallContext(context.Background(), &rawHeader, "eth_getBlockByNumber",
		hexutil.EncodeBig(big.NewInt(int64(blockNumber))), false); err != nil {
		return nil, err
	}

	var header *ethtypes.Header
	if err := json.Unmarshal(rawHeader, &header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal block header error: %w rawHeader: %s", err, rawHeader)
	}
	if header == nil {
		return nil, rpc.ErrNoResult
	}

	c.cache.Add(blockNumber, rawHeader)

	return rawHeader, nil
}
