package operations

import (
	"context"
	"fmt"

	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/consensus"
	"github.com/Lagrange-Labs/lagrange-node/governance"
	"github.com/Lagrange-Labs/lagrange-node/network"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/store"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/testutil"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

const (
	stakeAddr   = "0x98f07aB2d35638B79582b250C01444cEce0E517A"
	slasherAddr = "0x6Bf0fF4eBa00E3668c0241bb1C622CDBFE55bbE0"
)

// Manager is a struct for test operations.
type Manager struct {
	cfg     *config.Config
	chainID uint32
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
		if err := network.RunServer(&m.cfg.Server, m.Storage, state, m.chainID); err != nil {
			panic(err)
		}
	}()
}

// RegisterOperator registers a new operator.
func (m *Manager) RegisterOperator(privateKey string) {
	ethClient, err := ethclient.Dial(m.cfg.Client.RPCEndpoint)
	if err != nil {
		panic(fmt.Errorf("failed to connect to ethereum node: %w", err))
	}
	auth, err := utils.GetSigner(context.Background(), ethClient, privateKey)
	if err != nil {
		panic(fmt.Errorf("failed to get signer: %w", err))
	}

	if err := testutil.RegisterOperator(ethClient, auth, common.HexToAddress(stakeAddr), common.HexToAddress(slasherAddr)); err != nil {
		panic(fmt.Errorf("failed to register operator: %w", err))
	}
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
	clientCfg1.ECDSAPrivateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
	m.RegisterOperator(clientCfg1.ECDSAPrivateKey)
	m.RunClient(&clientCfg1)
	// client2
	clientCfg2 := m.cfg.Client
	clientCfg2.BLSPrivateKey = "0x475e7dc95f40ba8e5af29adb745ae3ac5d3404575b0f853c73ed1efa46943fc2"
	clientCfg2.ECDSAPrivateKey = "0x25f536330df3a72fa381bfb5ea5552b2731523f08580e7a0e2e69618a9643faa"
	m.RegisterOperator(clientCfg2.ECDSAPrivateKey)
	m.RunClient(&clientCfg2)
	// client3
	clientCfg3 := m.cfg.Client
	clientCfg3.BLSPrivateKey = "0x59ec5a675fa5a9805d791c58c97a3dcc0bc8def2029bd53aa33dc035f2b81404"
	clientCfg3.ECDSAPrivateKey = "0xc262364335471942e02e79d760d1f5c5ad7a34463303851cacdd15d72e68b228"
	m.RegisterOperator(clientCfg3.ECDSAPrivateKey)
	m.RunClient(&clientCfg3)
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
