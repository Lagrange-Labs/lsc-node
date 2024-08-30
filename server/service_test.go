package server

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/server/types"
	v2types "github.com/Lagrange-Labs/lagrange-node/server/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
)

type mockConsensus struct{}

func (m *mockConsensus) Start() {}

func (m *mockConsensus) IsStopped() bool {
	return false
}

func (m *mockConsensus) GetOpenBatch() *sequencerv2types.Batch {
	return &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{},
	}
}

func (m *mockConsensus) GetPrevBatch() *sequencerv2types.Batch {
	return &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{},
	}
}

func (m *mockConsensus) GetRoundInterval() time.Duration {
	return 0
}

func (m *mockConsensus) AddBatchCommit(commit *sequencerv2types.BlsSignature, stakeAddr, pubKey string) error {
	return nil
}

func (m *mockConsensus) CheckCommitteeMember(stakeAddr, pubKey string) (bool, error) {
	return true, nil
}

func (m *mockConsensus) CheckSignAddress(stakeAddr, signAddr string) bool {
	return true
}

func (m *mockConsensus) IsFinalized(batchkNumber uint64) bool {
	return true
}

func (m *mockConsensus) GetBLSScheme() crypto.BLSScheme {
	return nil
}

func newTestService() (*sequencerService, error) {
	storage, err := memdb.NewMemDB()
	if err != nil {
		return nil, err
	}

	storage.AddBlock(context.Background(), nil) //nolint:errcheck
	storage.AddBlock(context.Background(), nil) //nolint:errcheck
	storage.AddBlock(context.Background(), nil) //nolint:errcheck

	return &sequencerService{
		storage:   storage,
		consensus: &mockConsensus{},
		blsScheme: crypto.NewBLSScheme(crypto.BN254),
	}, nil
}

func TestBLSSignVerify(t *testing.T) {
	blsScheme := crypto.NewBLSScheme(crypto.BN254)
	priv, err := blsScheme.GenerateRandomKey()
	require.NoError(t, err)
	pub, err := blsScheme.GetPublicKey(priv, true, true)
	require.NoError(t, err)

	// JoinNetwork request sign
	req := &types.JoinNetworkRequest{
		StakeAddress: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		PublicKey:    core.Bytes2Hex(pub),
	}

	msg, err := proto.Marshal(req)
	require.NoError(t, err)

	sig, err := blsScheme.Sign(priv, msg, true)
	require.NoError(t, err)

	// Verify signature
	verified, err := blsScheme.VerifySignature(pub, msg, sig, true)
	require.NoError(t, err)
	require.True(t, verified)
}

func TestJoinNetwork(t *testing.T) {
	ctx := context.Background()
	peerCtx := peer.NewContext(ctx, &peer.Peer{
		Addr: &net.IPAddr{},
	})

	testCases := []struct {
		name      string
		ctx       context.Context
		stakeAdr  string
		pubKey    string
		signature string
		valid     bool
		wantErr   bool
	}{
		{"invalid signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b", "", false, true},
		{"wrong signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b", "a2e3cf2037699b3856c72af280ab8501878495dd81595128df23ba3de0e52fd9126c02b9262b871074f5a34495cd1a1c13cf3d27881ce9a8846463b7d30024c37861e0fa20418c186628f9b6565a116017f988f2d9ae058480fae910a4659bf0", false, true},
		{"invalid peer ctx", ctx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b", "9ce1d4e95d3191ef1e171838e5b451b849c3c4b3946fa6e87ed610f9160960300357bb907872325a9384e7625d3686f5580dd81218b44fe0d25dfdc48f6bee97", false, true},
		{"valid signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "8afdc78675918678650ad4cf045701e3535eb8b46e8b5425a99f2100a92ea06b", "9ce1d4e95d3191ef1e171838e5b451b849c3c4b3946fa6e87ed610f9160960300357bb907872325a9384e7625d3686f5580dd81218b44fe0d25dfdc48f6bee97", true, false},
	}

	service, err := newTestService()
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := service.JoinNetwork(tc.ctx, &v2types.JoinNetworkRequest{
				StakeAddress: tc.stakeAdr,
				PublicKey:    tc.pubKey,
				Signature:    tc.signature,
			})
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.valid {
					valid, err := ValidateToken(res.Token, tc.stakeAdr)
					require.NoError(t, err)
					require.True(t, valid)
				}
			}
		})
	}
}
