package operations

import (
	"os"
	"path/filepath"

	"github.com/Lagrange-Labs/lsc-node/config"
	"github.com/Lagrange-Labs/lsc-node/consensus"
	"github.com/Lagrange-Labs/lsc-node/core/crypto"
	"github.com/Lagrange-Labs/lsc-node/rpcclient"
	"github.com/Lagrange-Labs/lsc-node/sequencer"
	"github.com/Lagrange-Labs/lsc-node/server"
	"github.com/Lagrange-Labs/lsc-node/store"
	storetypes "github.com/Lagrange-Labs/lsc-node/store/types"
	"github.com/Lagrange-Labs/lsc-node/testutil"
)

// Manager is a struct for test operations.
type Manager struct {
	cfg       *config.Config
	chainInfo *consensus.ChainInfo
	sequencer *sequencer.Sequencer

	// Storage is a storage interface for test operations.
	Storage storetypes.Storage
}

// NewManager returns a operation manager.
func NewManager() (*Manager, error) {
	cfg, err := config.Default()
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpcclient.NewClient(cfg.Sequencer.Chain, &cfg.RpcClient, true)
	if err != nil {
		return nil, err
	}
	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return nil, err
	}
	store, err := store.NewStorage(&cfg.Store)
	return &Manager{
		cfg:     cfg,
		Storage: store,
		chainInfo: &consensus.ChainInfo{
			ChainID:            chainID,
			EthereumURL:        cfg.Sequencer.EthereumURL,
			CommitteeSCAddress: cfg.Sequencer.CommitteeSCAddress,
		},
	}, err
}

// RunServer runs a new server instance.
func (m *Manager) RunServer() {
	keystorePath := filepath.Join(os.TempDir(), "bls.json")
	if err := testutil.GenerateRandomKeystore(string(crypto.BN254), "password_localtest", keystorePath); err != nil {
		panic(err)
	}
	m.cfg.Consensus.ProposerBLSKeystorePath = keystorePath
	state := consensus.NewState(&m.cfg.Consensus, m.Storage, m.chainInfo)
	state.Start()
	go func() {
		if err := server.RunServer(&m.cfg.Server, m.Storage, state, m.chainInfo.ChainID); err != nil {
			panic(err)
		}
	}()
}

// RunSequencer runs a new sequencer instance.
func (m *Manager) RunSequencer(isGov bool) {
	var err error
	m.sequencer, err = sequencer.NewSequencer(&m.cfg.Sequencer, &m.cfg.RpcClient, m.Storage)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := m.sequencer.Start(); err != nil {
			panic(err)
		}
	}()
}

// GetChainID returns the chain id.
func (m *Manager) GetChainID() uint32 {
	if m.sequencer == nil {
		return 0
	}
	return m.sequencer.GetChainID()
}

// Close closes the manager.
func (m *Manager) Close() {
	if m.sequencer != nil {
		m.sequencer.Stop()
		m.sequencer = nil
	}
}

// GetRpcConfig returns the rpc config.
func (m *Manager) GetRpcConfig() *rpcclient.Config {
	return &m.cfg.RpcClient
}
