package mantle

import (
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

var _ types.RpcClient = (*Client)(nil)

// Client is a Mantle client.
type Client struct {
	evmclient.Client

	ethClient    *ethclient.Client
	batchStorage common.Address // Address of the L1BatchStorage contract
}

var (
	getL2BlockNumberABI abi.ABI
	abiInput            []byte
)

func init() {
	var err error
	getL2BlockNumberABI, err = abi.JSON(strings.NewReader(`[{"inputs":[],"name":"getL2StoredBlockNumber","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	abiInput, err = getL2BlockNumberABI.Pack("getL2StoredBlockNumber")
	if err != nil {
		panic(err)
	}
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:       *client,
		ethClient:    ethClient,
		batchStorage: common.HexToAddress(cfg.BatchStorageAddr),
	}, nil
}

// GetFinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	b, err := c.ethClient.BlockNumber(utils.GetContext())
	if err != nil {
		logger.Errorf("failed to get block number: %v", err)
		return 0, err
	}

	if b <= 64 {
		return 0, nil
	}

	// Get the L2 block number from the L1BatchStorage contract
	msg := ethereum.CallMsg{
		To:   &c.batchStorage,
		Data: abiInput,
	}

	result, err := c.ethClient.CallContract(utils.GetContext(), msg, big.NewInt(int64(b-64)))
	if err != nil {
		if strings.Contains(err.Error(), "missing trie node") {
			// TODO: This is a temporary workaround for the missing trie node error.
			// It means the dedicated RPC node is not fully synced yet.
			logger.Infof("Missing trie node error: %v", err)
			return math.MaxUint64, nil
		}
		logger.Errorf("failed to call L1BatchStorage contract: %v", err)
		return 0, err
	}

	var blockNumber *big.Int
	err = getL2BlockNumberABI.UnpackIntoInterface(&blockNumber, "getL2StoredBlockNumber", result)

	return blockNumber.Uint64(), err
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber uint64) bool {
	return true
}

// NextBatch returns the next batch after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	return nil, types.ErrNoResult
}
