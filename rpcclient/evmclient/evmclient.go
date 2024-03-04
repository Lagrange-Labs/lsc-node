package evmclient

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

type Client struct {
	ethClient *ethclient.Client
	rpcClient *rpc.Client
}

var _ types.RpcClient = (*Client)(nil)

// NewClient creates a new EvmClient instance.
func NewClient(rpcURL string) (*Client, error) {
	rpcClient, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcClient: rpcClient,
		ethClient: ethclient.NewClient(rpcClient),
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
	result := &struct {
		Hash common.Hash `json:"hash"`
	}{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.rpcClient.CallContext(ctx, &result, "eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(blockNumber))), false)
	if err == nil && result == nil {
		return "", types.ErrBlockNotFound
	} else if err != nil {
		return "", err
	}

	return result.Hash.Hex(), nil
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
	var result *ethtypes.Header

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.rpcClient.CallContext(ctx, &result, "eth_getBlockByNumber", "finalized", false)
	if err != nil || result == nil || result.Number == nil {
		// TODO: finalized block API is not supported
		return math.MaxUint64, nil
	}

	return result.Number.Uint64(), nil
}
