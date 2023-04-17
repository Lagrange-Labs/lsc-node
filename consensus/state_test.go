package consensus

import (
	"context"
	"sync"
	"testing"
	"time"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/store"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func createTestState(t *testing.T) (*State, chan *networktypes.CommitBlockRequest) {
	_, pubKey := utils.RandomBlsKey()
	cfg := &Config{
		ProposerPubKey: pubKey,
		RoundLimit:     utils.TimeDuration(5 * time.Second),
		RoundInterval:  utils.TimeDuration(2 * time.Second),
	}

	memDB, err := store.NewMemDB()
	require.NoError(t, err)
	require.NoError(t, memDB.AddBlock(context.Background(), nil))

	chCommit := make(chan *networktypes.CommitBlockRequest)
	return NewState(cfg, memDB, chCommit), chCommit
}

func TestState_OnStart(t *testing.T) {
	s, _ := createTestState(t)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := s.OnStart()
		require.NoError(t, err)
		t.Log(err)
	}()

	time.Sleep(1 * time.Second)

	s.OnStop()

	wg.Wait()
}
