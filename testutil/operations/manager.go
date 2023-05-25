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
	"github.com/urfave/cli/v2"
)

// Manager is a struct for test operations.
type Manager struct {
	cfg     *config.Config
	chainID int32
	// Storage is a storage interface for test operations.
	Storage storetypes.Storage
}

// NewManager returns a operation manager.
func NewManager() (*Manager, error) {
	cfg, err := config.Load(&cli.Context{})
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpcclient.CreateRPCClient(cfg.Sequencer.Chain, cfg.Sequencer.RPCURL)
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
		if err := network.RunServer(&m.cfg.Server, m.Storage, state); err != nil {
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
	clientCfg := m.cfg.Client
	clientCfg.BLSPrivateKey = "0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f654a"
	clientCfg.ECDSAPrivateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
	m.RunClient(&clientCfg)
	// client2
	clientCfg = m.cfg.Client
	clientCfg.BLSPrivateKey = "0xca056e1830045cd5aa844e02e493e27b8c454b3059315b8ee34d42414141247d"
	clientCfg.ECDSAPrivateKey = "0x25f536330df3a72fa381bfb5ea5552b2731523f08580e7a0e2e69618a9643faa"
	m.RunClient(&clientCfg)
	// client3
	clientCfg = m.cfg.Client
	clientCfg.BLSPrivateKey = "0xa339d4976df7cc511ad3d31fb28407e0b92a3f876cc8d81434abfe27e09c0275"
	clientCfg.ECDSAPrivateKey = "0xc262364335471942e02e79d760d1f5c5ad7a34463303851cacdd15d72e68b228"
	m.RunClient(&clientCfg)
}

// RunSequencer runs a new sequencer instance.
func (m *Manager) RunSequencer() {
	sequencer, err := sequencer.NewSequencer(&m.cfg.Sequencer, m.Storage)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := sequencer.Start(); err != nil {
			panic(err)
		}
	}()

	gov, err := governance.NewGovernance(&m.cfg.Governance, m.Storage)
	if err != nil {
		panic(err)
	}
	go gov.Start()
}
