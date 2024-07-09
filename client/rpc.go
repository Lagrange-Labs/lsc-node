package client

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	committeeCacheSize = 10
)

type rpcAdapter struct {
	client         rpctypes.RpcClient
	committeeSC    *committee.Committee
	committeeCache *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]

	chainID            uint32
	genesisBlockNumber uint64
}

// NewRpcAdapter creates a new rpc adapter.
func NewRpcAdapter(rpcCfg *rpcclient.Config, cfg *Config) (*rpcAdapter, error) {
	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create the rpc client: %v, please check the chain name, the chain name should look like 'optimism', 'base'", err)
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return nil, fmt.Errorf("failed to get the chain ID: %v", err)
	}

	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create the ethereum client: %v", err)
	}
	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create the committee smart contract: %v", err)
	}
	params, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		logger.Fatalf("failed to get the committee params: %v", err)
	}

	return &rpcAdapter{
		client:             rpcClient,
		committeeSC:        committeeSC,
		committeeCache:     lru.NewCache[uint64, *committee.ILagrangeCommitteeCommitteeData](committeeCacheSize),
		chainID:            chainID,
		genesisBlockNumber: uint64(params.GenesisBlock.Int64() - params.L1Bias.Int64()),
	}, nil
}

func (r *rpcAdapter) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) {
	if committeeData, ok := r.committeeCache.Get(blockNumber); ok {
		return committeeData, nil
	}

	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_committee")

	committeeData, err := r.committeeSC.GetCommittee(nil, r.chainID, big.NewInt(int64(blockNumber)))
	if err != nil || committeeData.LeafCount == 0 {
		return nil, fmt.Errorf("failed to get the committee data %+v: %v", committeeData, err)
	}
	r.committeeCache.Add(blockNumber, &committeeData)

	return &committeeData, nil
}

func (r *rpcAdapter) setBeginBlockNumber(blockNumber uint64) bool {
	return r.client.SetBeginBlockNumber(blockNumber)
}

func (r *rpcAdapter) nextBatch() (*sequencerv2types.BatchHeader, error) {
	return r.client.NextBatch()
}
