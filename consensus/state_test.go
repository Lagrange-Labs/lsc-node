package consensus

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
	"github.com/Lagrange-Labs/lagrange-node/testutil"
)

func createTestState(t *testing.T) (*State, chan *sequencertypes.BlsSignature) {
	keystorePassword := "password"
	keystorePath := filepath.Join(t.TempDir(), "bls.json")
	err := testutil.GenerateRandomKeystore(string(crypto.BN254), keystorePassword, keystorePath)
	require.NoError(t, err)
	cfg := &Config{
		ProposerBLSKeystorePath:     keystorePath,
		ProposerBLSKeystorePassword: keystorePassword,
		RoundLimit:                  core.TimeDuration(5 * time.Second),
		RoundInterval:               core.TimeDuration(2 * time.Second),
		BLSCurve:                    string(crypto.BN254),
	}

	memDB, err := memdb.NewMemDB()
	require.NoError(t, err)
	require.NoError(t, memDB.AddBlock(context.Background(), nil))

	chCommit := make(chan *sequencertypes.BlsSignature)
	return NewState(cfg, memDB, &ChainInfo{
		ChainID:            1,
		EthereumURL:        "http://localhost:8545",
		CommitteeSCAddress: "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06",
	}), chCommit
}

func TestState_OnStart(t *testing.T) {
	s, _ := createTestState(t)

	s.Start()

	time.Sleep(1 * time.Second)

	s.Stop()
}
