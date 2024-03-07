package governance

import (
	"context"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/config"
	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/governance/types"
	"github.com/Lagrange-Labs/lagrange-node/store"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

const (
	privateKey = "0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429"
)

var (
	committeeAddr = "0x"
)

func init() {
	cfg, err := config.Default()
	if err != nil {
		panic(err)
	}
	committeeAddr = cfg.Governance.CommitteeSCAddress
}

func createTestGovernance(t *testing.T) (storetypes.Storage, *Governance, *ethclient.Client, *bind.TransactOpts) {
	storeCfg := store.Config{
		BackendType: "mongodb",
		DBPath:      "mongodb://127.0.0.1:27017",
	}
	storage, err := store.NewStorage(&storeCfg)
	require.NoError(t, err)

	govCfg := types.Config{
		EthereumURL:          "http://localhost:8545",
		PrivateKey:           "0x3e17bc938ec10c865fc4e2d049902716dc0712b5b0e688b7183c16807234a84c",
		CommitteeSCAddress:   committeeAddr,
		StakingCheckInterval: utils.TimeDuration(time.Second * 1),
	}
	client, err := ethclient.Dial(govCfg.EthereumURL)
	require.NoError(t, err)
	auth, err := utils.GetSigner(context.Background(), client, privateKey)
	require.NoError(t, err)
	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)
	gov, err := NewGovernance(&govCfg, crypto.BN254, uint32(chainID.Int64()), storage)
	require.NoError(t, err)
	return storage, gov, client, auth
}

func TestUploadEvidence(t *testing.T) {
	// create the test governance
	storage, _, _, auth := createTestGovernance(t)
	// add the evidence to the storage
	blockHash := common.HexToHash(utils.RandomHex(32))
	committeeRoot := common.HexToHash(utils.RandomHex(32))
	evidence := &contypes.Evidence{
		Operator:             auth.From.Hex(),
		BlockHash:            blockHash,
		CurrentCommitteeRoot: committeeRoot,
		NextCommitteeRoot:    committeeRoot,
		BlockNumber:          1,
		L1BlockNumber:        1,
		BlockSignature:       common.FromHex(utils.RandomHex(32)),
		CommitSignature:      common.FromHex(utils.RandomHex(32)),
		ChainID:              1,
	}
	require.NoError(t, storage.AddEvidences(context.Background(), []*contypes.Evidence{evidence}))
	// check the evidence status
	evidences, err := storage.GetEvidences(context.Background(), 1, 1, 1, 1, 0)
	require.NoError(t, err)
	require.Equal(t, 1, len(evidences))
}
