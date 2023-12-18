package operations

import (
	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/consensus"
	"github.com/Lagrange-Labs/lagrange-node/governance"
	"github.com/Lagrange-Labs/lagrange-node/network"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/store"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

// Manager is a struct for test operations.
type Manager struct {
	cfg       *config.Config
	chainID   uint32
	sequencer *sequencer.Sequencer
	gov       *governance.Governance
	// Storage is a storage interface for test operations.
	Storage storetypes.Storage
}

// NewManager returns a operation manager.
func NewManager() (*Manager, error) {
	cfg, err := config.Default()
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpcclient.NewClient(cfg.Sequencer.Chain, cfg.Sequencer.RPCURL, cfg.Sequencer.EthURL, cfg.Sequencer.BatchStorageAddr)
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
		chainID: chainID,
		Storage: store,
	}, err
}

// RunServer runs a new server instance.
func (m *Manager) RunServer() {
	state := consensus.NewState(&m.cfg.Consensus, m.Storage, m.chainID)
	go state.OnStart()
	go func() {
		if err := network.RunServer(&m.cfg.Server, m.Storage, state, m.chainID); err != nil {
			panic(err)
		}
	}()
}

// RunClient runs a new client instance.
func (m *Manager) RunClient(clientCfg *network.ClientConfig) {
	client, err := network.NewClient(clientCfg)
	if err != nil {
		panic(err)
	}
	go client.Start()
}

// RunClients runs several client instances.
func (m *Manager) RunClients() {
	// client1
	clientCfg1 := m.cfg.Client
	clientCfg1.BLSPrivateKey = "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a"
	clientCfg1.ECDSAPrivateKey = "0x220ecb0a36b61b15a3af292c4520a528395edc51c8d41db30c74382a4af4328d"
	m.RunClient(&clientCfg1)
	// client2
	clientCfg2 := m.cfg.Client
	clientCfg2.BLSPrivateKey = "0x475e7dc95f40ba8e5af29adb745ae3ac5d3404575b0f853c73ed1efa46943fc2"
	clientCfg2.ECDSAPrivateKey = "0x25f536330df3a72fa381bfb5ea5552b2731523f08580e7a0e2e69618a9643faa"
	m.RunClient(&clientCfg2)
	// client3
	clientCfg3 := m.cfg.Client
	clientCfg3.BLSPrivateKey = "0x59ec5a675fa5a9805d791c58c97a3dcc0bc8def2029bd53aa33dc035f2b81404"
	clientCfg3.ECDSAPrivateKey = "0xc262364335471942e02e79d760d1f5c5ad7a34463303851cacdd15d72e68b228"
	m.RunClient(&clientCfg3)
}

// RunSequencer runs a new sequencer instance.
func (m *Manager) RunSequencer() {
	var err error
	m.sequencer, err = sequencer.NewSequencer(&m.cfg.Sequencer, m.Storage)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := m.sequencer.Start(); err != nil {
			panic(err)
		}
	}()

	m.gov, err = governance.NewGovernance(&m.cfg.Governance, m.sequencer.GetChainID(), m.Storage)
	if err != nil {
		panic(err)
	}
	go m.gov.Start()
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
	if m.gov != nil {
		m.gov.Stop()
		m.gov = nil
	}
}
