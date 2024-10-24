package client

import (
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lsc-node/core/telemetry"
	"github.com/Lagrange-Labs/lsc-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lsc-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lsc-node/sequencer/types/v2"
	serverv2types "github.com/Lagrange-Labs/lsc-node/server/types/v2"
	"github.com/ethereum/go-ethereum/common"
)

var _ AdapterCaller = (*RpcAdapter)(nil)

// RpcAdapter is the adapter for the RPC client.
type RpcAdapter struct {
	client       rpctypes.RpcClient
	chainID      uint32
	cache        [2]*sequencerv2types.BatchHeader
	chNodeStatus chan<- StatusMessage
}

// newRpcAdapter creates a new rpc adapter.
func newRpcAdapter(rpcCfg *rpcclient.Config, cfg *Config, accountID string, chNodeStatus chan<- StatusMessage) (*RpcAdapter, uint32, error) {
	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg, true)
	if err != nil {
		chNodeStatus <- StatusMessage{
			NodeStatus: serverv2types.ClientNodeStatus_L2_RPC_ISSUE,
			Message:    fmt.Sprintf("failed to create the rpc client: %v", cfg.Chain),
		}
		return nil, 0, fmt.Errorf("failed to create the rpc client: %v, please check the chain name, the chain name should look like 'optimism', 'base'", err)
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		chNodeStatus <- StatusMessage{
			NodeStatus: serverv2types.ClientNodeStatus_L2_RPC_ISSUE,
			Message:    fmt.Sprintf("failed to get the chain ID for %s", cfg.Chain),
		}
		return nil, 0, fmt.Errorf("failed to get the chain ID: %v", err)
	}

	return &RpcAdapter{
		client:       rpcClient,
		chainID:      chainID,
		chNodeStatus: chNodeStatus,
	}, chainID, nil
}

func (r *RpcAdapter) checkCache(l1BlockNumber uint64, l1TxIndex uint32) bool {
	if r.cache[1] == nil {
		return false
	}

	if r.cache[1].L1BlockNumber == l1BlockNumber && r.cache[1].L1TxIndex == l1TxIndex {
		return true
	}

	return false
}

func (r *RpcAdapter) updateCache(batchHeader *sequencerv2types.BatchHeader) {
	if r.cache[0] == nil {
		r.cache[0] = batchHeader
		return
	}
	if r.cache[1] == nil {
		r.cache[1] = batchHeader
		return
	}
	r.cache[0] = r.cache[1]
	r.cache[1] = batchHeader
}

// GetPrevBatchL1Number gets the previous batch L1 number from the database.
func (r *RpcAdapter) GetPrevBatchL1Number(l1BlockNumber uint64, l1TxIndex uint32) (uint64, error) {
	if !r.checkCache(l1BlockNumber, l1TxIndex) {
		r.chNodeStatus <- StatusMessage{
			NodeStatus: serverv2types.ClientNodeStatus_INTERNAL_ISSUE,
			Message:    fmt.Sprintf("no previous batch for L1BlockNumber: %d L1TxIndex: %d", l1BlockNumber, l1TxIndex),
		}
		return 0, fmt.Errorf("no previous batch for L1BlockNumber: %d L1TxIndex: %d", l1BlockNumber, l1TxIndex)
	}
	return r.cache[0].L1BlockNumber, nil
}

// GetBatchHeader gets the batch header from the database.
func (r *RpcAdapter) GetBatchHeader(l1BlockNumber uint64, txHash string, l1TxIndex uint32) (*sequencerv2types.BatchHeader, error) {
	if r.checkCache(l1BlockNumber, l1TxIndex) {
		return r.cache[1], nil
	}

	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_batch_header")

	batchHeader, err := r.client.GetL2BatchHeader(l1BlockNumber, txHash)
	if err != nil {
		r.chNodeStatus <- StatusMessage{
			NodeStatus: serverv2types.ClientNodeStatus_L2_RPC_ISSUE,
			Message:    fmt.Sprintf("failed to get the batch header for L1BlockNumber: %d TxHash: %v err: %v", l1BlockNumber, txHash, err),
		}
		return nil, err
	}

	r.updateCache(batchHeader)
	return batchHeader, nil
}

// VerifyBatchHeader verifies if the batch exists for the given L1 and L2 block number.
func (r *RpcAdapter) VerifyBatchHeader(l1BlockNumber, l2BlockNumber uint64) error {
	for i := 0; i < 2; i++ {
		if r.cache[i] != nil && r.cache[i].L1BlockNumber == l1BlockNumber && r.cache[i].L2FromBlockNumber == l2BlockNumber {
			// init cache
			r.cache[0] = r.cache[i]
			r.cache[1] = nil
			return nil
		}
	}

	batchHeader, err := r.client.VerifyBatchHeader(l1BlockNumber, l2BlockNumber)
	if err != nil {
		r.chNodeStatus <- StatusMessage{
			NodeStatus: serverv2types.ClientNodeStatus_L2_RPC_ISSUE,
			Message:    fmt.Sprintf("failed to verify the batch header for L1BlockNumber: %d L2BlockNumber: %d err: %v", l1BlockNumber, l2BlockNumber, err),
		}
		return err
	}

	r.cache[0] = batchHeader
	r.cache[1] = nil

	return nil
}

// GetBlockHash implements the Adapter interface.
func (r *RpcAdapter) GetBlockHash(rlpHeader []byte) (common.Hash, common.Hash, error) {
	return r.client.GetBlockHashFromRLPHeader(rlpHeader)
}
