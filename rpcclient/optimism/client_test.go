package optimism

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var _ types.EvmClient = (*MockEvmClient)(nil)

type MockEvmClient struct {
	triggerTimeErr time.Time
}

func NewMockEvmClient(triggerDur time.Duration) *MockEvmClient {
	return &MockEvmClient{
		triggerTimeErr: time.Now().Add(triggerDur),
	}
}

func (c *MockEvmClient) GetChainID() (uint32, error) {
	if time.Now().After(c.triggerTimeErr) {
		return 0, errors.New("error")
	}
	return 0, nil
}

func (c *MockEvmClient) GetBlockHashByNumber(blockNumber uint64) (common.Hash, error) {
	if time.Now().After(c.triggerTimeErr) {
		return common.Hash{}, errors.New("error")
	}
	return common.Hash{}, nil
}

func (c *MockEvmClient) GetBlockNumberByHash(blockHash common.Hash) (uint64, error) {
	if time.Now().After(c.triggerTimeErr) {
		return 0, errors.New("error")
	}
	return 0, nil
}

func (c *MockEvmClient) GetBlockNumberByTxHash(txHash common.Hash) (uint64, error) {
	if time.Now().After(c.triggerTimeErr) {
		return 0, errors.New("error")
	}
	return 0, nil
}

func (c *MockEvmClient) GetFinalizedBlockNumber() (uint64, error) {
	if time.Now().After(c.triggerTimeErr) {
		return 0, errors.New("error")
	}
	return math.MaxUint64, nil
}

func (c *MockEvmClient) GetBlockHashesByRange(startBlockNumber, endBlockNumber uint64) ([]common.Hash, error) {
	if time.Now().After(c.triggerTimeErr) {
		return nil, errors.New("error")
	}
	return nil, nil
}

func TestErrorHandling(t *testing.T) {
	cfg := &Config{
		RPCURL:             "http://localhost:8545",
		L1RPCURL:           "http://localhost:8545",
		BeaconURL:          "http://localhost:8545",
		BatchInbox:         common.Address{}.Hex(),
		BatchSender:        common.Address{}.Hex(),
		ConcurrentFetchers: 4,
	}
	client, err := NewClient(cfg)
	require.NoError(t, err)

	client.fetcher.l2Client = NewMockEvmClient(1 * time.Second)

	client.SetBeginBlockNumber(10, 10)
	time.Sleep(2 * time.Second)
	_, err = client.NextBatch()
	require.Error(t, err)
	// check if able to restart
	client.SetBeginBlockNumber(10, 10)
}
