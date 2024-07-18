package evmclient

import (
	"bytes"
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
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CacheSize = 2048

var _ types.EvmClient = (*Client)(nil)

// Client is an EVM client.
type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	rpcURLs   []string
	index     int

	cache *lru.Cache[uint64, json.RawMessage] // block number -> raw header
}

// NewClient creates a new EvmClient instance.
func NewClient(rpcURLs []string) (*Client, error) {
	c := &Client{
		rpcURLs: rpcURLs,
		index:   len(rpcURLs) - 1,
		cache:   lru.NewCache[uint64, json.RawMessage](CacheSize),
	}

	if err := c.switchRPCURL(); err != nil {
		return nil, err
	}

	return c, nil
}

// GetEthClient returns the ethclient.Client.
func (c *Client) GetEthClient() *ethclient.Client {
	return c.ethClient
}

// switchRPCURL switches the RPC URL.
func (c *Client) switchRPCURL() error {
	if c.rpcClient != nil {
		c.rpcClient.Close()
	}
	for i := 0; i < len(c.rpcURLs); i++ {
		c.index = (c.index + 1) % len(c.rpcURLs)
		rpcClient, err := rpc.DialContext(utils.GetContext(), c.rpcURLs[c.index])
		if err != nil {
			logger.Errorf("failed to dial the RPC URL error: %w", err)
			continue
		}
		c.rpcClient = rpcClient
		c.ethClient = ethclient.NewClient(rpcClient)
		return nil
	}

	return fmt.Errorf("failed to switch the RPC URL")
}

// switchCall tries to switch the RPC URL and call the given method.
func (c *Client) switchCall(fn func() (interface{}, error)) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	for i := 0; i <= 1; i++ {
		if res, err = fn(); err == nil {
			return res, nil
		}
		if err := c.switchRPCURL(); err != nil {
			return nil, err
		}
	}

	return nil, fmt.Errorf("failed to call the method error: %w", err)
}

// GetCurrentBlockNumber returns the current block number.
func (c *Client) GetCurrentBlockNumber() (uint64, error) {
	fn := func() (interface{}, error) {
		return c.ethClient.HeaderByNumber(utils.GetContext(), nil)
	}
	header, err := c.switchCall(fn)
	if err != nil {
		return 0, err
	}

	return header.(*ethtypes.Header).Number.Uint64(), nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *Client) GetBlockHashByNumber(blockNumber uint64) (common.Hash, error) {
	rawHeader, err := c.GetRawHeaderByNumber(blockNumber)
	if errors.Is(err, rpc.ErrNoResult) {
		return common.Hash{}, types.ErrBlockNotFound
	}
	if err != nil {
		logger.Errorf("failed to get the raw header error: %v", err)
		return common.Hash{}, err
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
	fn := func() (interface{}, error) {
		return c.ethClient.HeaderByHash(utils.GetContext(), blockHash)
	}
	header, err := c.switchCall(fn)
	if err != nil {
		return 0, err
	}
	return header.(*ethtypes.Header).Number.Uint64(), nil
}

// GetBlockNumberByTxHash returns the block number by the given transaction hash.
func (c *Client) GetBlockNumberByTxHash(txHash common.Hash) (uint64, error) {
	fn := func() (interface{}, error) {
		return c.ethClient.TransactionReceipt(utils.GetContext(), txHash)
	}
	receipt, err := c.switchCall(fn)
	if err != nil {
		return 0, err
	}
	return receipt.(*ethtypes.Receipt).BlockNumber.Uint64(), nil
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() (uint32, error) {
	fn := func() (interface{}, error) {
		return c.ethClient.ChainID(utils.GetContext())
	}
	chainID, err := c.switchCall(fn)
	if err != nil {
		return 0, err
	}
	return uint32(chainID.(*big.Int).Int64()), nil
}

// GetBlockHeadersByRange returns the block headers by the given block number range.
// The range is [start, end).
func (c *Client) GetBlockHeadersByRange(start, end uint64) ([]sequencerv2types.BlockHeader, error) {
	fn := func() (interface{}, error) {
		batchElems := make([]rpc.BatchElem, 0, int(end-start))
		for i := start; i < end; i++ {
			batchElems = append(batchElems, rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{hexutil.EncodeBig(big.NewInt(int64(i))), false},
				Result: &json.RawMessage{},
			})
		}
		err := c.rpcClient.BatchCallContext(utils.GetContext(), batchElems)
		return batchElems, err
	}
	res, err := c.switchCall(fn)
	if err != nil {
		return nil, err
	}
	batchElems := res.([]rpc.BatchElem)
	blockHeaders := make([]sequencerv2types.BlockHeader, len(batchElems))
	var buffer bytes.Buffer
	for i, batch := range batchElems {
		if batch.Error != nil {
			return nil, batch.Error
		}
		rawHeader := batch.Result.(*json.RawMessage)
		c.cache.Add(start+uint64(i), *rawHeader)
		blockHash, err := getHashFromRawHeader(*rawHeader)
		if err != nil {
			return nil, err
		}
		header := &ethtypes.Header{}
		if err := json.Unmarshal(*rawHeader, &header); err != nil {
			return nil, err
		}
		buffer.Reset()
		if err := rlp.Encode(&buffer, header); err != nil {
			return nil, err
		}
		blockHeaders[i] = sequencerv2types.BlockHeader{
			BlockNumber: start + uint64(i),
			BlockHash:   blockHash.Hex(),
			BlockRlp:    common.Bytes2Hex(buffer.Bytes()),
		}
	}
	return blockHeaders, nil
}

// GetFinalizedBlockNumber returns the finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	fn := func(tag string) func() (interface{}, error) {
		return func() (interface{}, error) {
			var header *ethtypes.Header
			err := c.rpcClient.CallContext(utils.GetContext(), &header, "eth_getBlockByNumber", tag, false)
			return header, err
		}
	}
	res, err := c.switchCall(fn("finalized"))
	if err != nil {
		if strings.Contains(err.Error(), "'finalized' tag not supported on pre-merge network") {
			res, err := c.switchCall(fn("latest"))
			if err != nil {
				logger.Errorf("failed to get latest block number error: %v", err)
				return 0, err
			}
			return res.(*ethtypes.Header).Number.Uint64(), nil
		}
		logger.Errorf("failed to get finalized block number error: %v", err)
		return 0, err
	}
	return res.(*ethtypes.Header).Number.Uint64(), nil
}

// GetRawHeaderByNumber returns the raw message of block header by the given block number.
func (c *Client) GetRawHeaderByNumber(blockNumber uint64) (json.RawMessage, error) {
	if raw, ok := c.cache.Get(blockNumber); ok {
		return raw, nil
	}

	fn := func() (interface{}, error) {
		var rawHeader json.RawMessage
		err := c.rpcClient.CallContext(utils.GetContext(), &rawHeader, "eth_getBlockByNumber",
			hexutil.EncodeBig(big.NewInt(int64(blockNumber))), false)
		return rawHeader, err
	}
	res, err := c.switchCall(fn)
	if err != nil {
		if strings.Contains(err.Error(), rpc.ErrNoResult.Error()) {
			return nil, types.ErrNoResult
		}
		return nil, err
	}
	rawHeader := res.(json.RawMessage)
	if len(rawHeader) < 5 { // to detect empty response and "null" response
		return nil, types.ErrNoResult
	}

	c.cache.Add(blockNumber, rawHeader)

	return rawHeader, nil
}

// GetBlockByNumber returns the block header by the given block number.
func (c *Client) GetBlockByNumber(blockNumber uint64) (*ethtypes.Block, error) {
	fn := func() (interface{}, error) {
		return c.ethClient.BlockByNumber(utils.GetContext(), big.NewInt(int64(blockNumber)))
	}
	b, err := c.switchCall(fn)
	if err != nil {
		return nil, err
	}
	return b.(*ethtypes.Block), nil
}

// GetBlockHashFromRLPHeader returns the block hash and the parent hash from the rlp encoded block header.
func (c *Client) GetBlockHashFromRLPHeader(rlpHeader []byte) (common.Hash, common.Hash, error) {
	header := &ethtypes.Header{}
	if err := rlp.Decode(bytes.NewReader(rlpHeader), header); err != nil {
		return common.Hash{}, common.Hash{}, err
	}

	return header.Hash(), header.ParentHash, nil
}
