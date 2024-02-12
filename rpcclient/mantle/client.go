package mantle

import (
	"context"
	"math"
	"math/big"
	"strings"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
func NewClient(rpcURL, l1RpcURL string, batchStorageAddr string) (*Client, error) {
	client, err := evmclient.NewClient(rpcURL)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(l1RpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:       *client,
		ethClient:    ethClient,
		batchStorage: common.HexToAddress(batchStorageAddr),
	}, nil
}

// GetL2FinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetL2FinalizedBlockNumber() (uint64, error) {
	b, err := c.ethClient.BlockNumber(context.Background())
	if err != nil {
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

	result, err := c.ethClient.CallContract(context.Background(), msg, big.NewInt(int64(b-64)))
	if err != nil {
		if strings.Contains(err.Error(), "missing trie node") {
			// TODO: This is a temporary workaround for the missing trie node error.
			// It means the dedicated RPC node is not fully synced yet.
			logger.Infof("Missing trie node error: %v", err)
			return math.MaxUint64, nil
		}
		// TODO: This is a temporary workaround for the storage contract not being deployed yet.
		curL2Number, err := c.GetCurrentBlockNumber()
		if err != nil {
			return 0, err
		}
		return curL2Number - 1000, nil
	}

	var blockNumber *big.Int
	err = getL2BlockNumberABI.UnpackIntoInterface(&blockNumber, "getL2StoredBlockNumber", result)

	return blockNumber.Uint64(), err
}
