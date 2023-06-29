package governance

import (
	"context"
	"testing"
	"time"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/store"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

const (
	stakeAddr  = "0xf32358f5C8FFfCF1a7bDb58b270a082abb7Ba1A6"
	privateKEy = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
	chainID    = 1337
)

func createTestGovernance(t *testing.T) (storetypes.Storage, *Governance, *ethclient.Client, *bind.TransactOpts) {
	storeCfg := store.Config{
		BackendType: "mongodb",
		DBPath:      "mongodb://localhost:27017",
	}
	storage, err := store.NewStorage(&storeCfg)
	require.NoError(t, err)

	govCfg := Config{
		EthereumURL:            "http://localhost:8545",
		PrivateKey:             "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c",
		StakingSCAddress:       stakeAddr,
		CommitteeSCAddress:     "0x11b59cE6E2b4509218bf45AF3582dC7E2a1e8a57",
		StakingCheckInterval:   utils.TimeDuration(time.Second * 1),
		EvidenceUploadInterval: utils.TimeDuration(time.Second * 1),
	}
	client, err := ethclient.Dial(govCfg.EthereumURL)
	require.NoError(t, err)
	auth, err := utils.GetSigner(context.Background(), client, privateKEy)
	require.NoError(t, err)
	gov, err := NewGovernance(&govCfg, chainID, storage)
	require.NoError(t, err)
	return storage, gov, client, auth
}

func TestUpdateNodeStatus(t *testing.T) {
	// create the test governance and register the operator
	storage, gov, _, auth := createTestGovernance(t)
	// join the network
	clientNode := networktypes.ClientNode{
		PublicKey:    "0x123",
		IPAddress:    "127.0.0.1",
		StakeAddress: auth.From.Hex(),
		ChainID:      chainID,
	}
	require.NoError(t, storage.AddNode(context.Background(), &clientNode))
	// update the node status
	time.Sleep(time.Second * 3)
	require.NoError(t, gov.updateNodeStatuses())
	// check the node status
	node, err := storage.GetNodeByStakeAddr(context.Background(), clientNode.StakeAddress)
	require.NoError(t, err)
	require.Equal(t, auth.From.Hex(), node.StakeAddress)
	require.Equal(t, networktypes.NodeRegistered, node.Status)
	// update the node status
	clientNode.Status = networktypes.NodeSlashed
	require.NoError(t, storage.UpdateNode(context.Background(), &clientNode))
}

func TestUploadEvidence(t *testing.T) {
	// create the test governance
	storage, _, _, auth := createTestGovernance(t)
	// add the evidence to the storage
	blockHash := common.HexToHash(utils.RandomHex(32))
	committeeRoot := common.HexToHash(utils.RandomHex(32))
	evidence := &contypes.Evidence{
		Operator:                    auth.From.Hex(),
		BlockHash:                   blockHash,
		CorrectBlockHash:            blockHash,
		CurrentCommitteeRoot:        committeeRoot,
		CorrectCurrentCommitteeRoot: committeeRoot,
		NextCommitteeRoot:           committeeRoot,
		CorrectNextCommitteeRoot:    committeeRoot,
		BlockNumber:                 1,
		EpochBlockNumber:            1,
		BlockSignature:              common.FromHex(utils.RandomHex(32)),
		CommitSignature:             common.FromHex(utils.RandomHex(32)),
		ChainID:                     1,
	}
	require.NoError(t, storage.AddEvidences(context.Background(), []*contypes.Evidence{evidence}))
	// check the evidence status
	evidences, err := storage.GetEvidences(context.Background())
	require.NoError(t, err)
	require.Equal(t, 1, len(evidences))
	// upload the evidence
	// require.NoError(t, gov.uploadEvidences())
	// // check the evidence status
	// evidences, err = storage.GetEvidences(context.Background())
	// require.NoError(t, err)
	// require.Equal(t, 0, len(evidences))
}
