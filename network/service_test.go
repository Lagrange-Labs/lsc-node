package network

import (
	"context"
	"net"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func newTestService() (*sequencerService, error) {
	storage, err := store.NewMemDB()
	if err != nil {
		return nil, err
	}

	storage.AddBlock(context.Background(), nil)
	storage.AddBlock(context.Background(), nil)
	storage.AddBlock(context.Background(), nil)

	return &sequencerService{
		storage:   storage,
		threshold: 1,
	}, nil
}

func TestBLSSign(t *testing.T) {
	privKey := "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	priv := new(bls.SecretKey)
	err := priv.Unmarshal(common.FromHex(privKey))
	require.NoError(t, err)

	pub := priv.GetPublicKey().Serialize()
	t.Log(common.Bytes2Hex(pub[:]))

	// JoinNetwork request sign
	req := &types.JoinNetworkRequest{
		StakeAddress: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		PublicKey:    "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4",
	}

	msg, err := proto.Marshal(req)
	require.NoError(t, err)

	sig, err := priv.Sign(msg)
	require.NoError(t, err)
	sigMsg := sig.Serialize()
	t.Log(common.Bytes2Hex(sigMsg[:]))
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
		{"invalid signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4", "", false, true},
		{"wrong signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4", "a2e3cf2037699b3856c72af280ab8501878495dd81595128df23ba3de0e52fd9126c02b9262b871074f5a34495cd1a1c13cf3d27881ce9a8846463b7d30024c37861e0fa20418c186628f9b6565a116017f988f2d9ae058480fae910a4659bf0", false, true},
		{"invalid peer ctx", ctx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4", "a2e3cf2037699b3856c72af280ab8501878495dd81595128df23ba3de0e52fd9126c02b9262b871074f5a34495cd1a1c13cf3d27881ce9a8846463b7d30024c37861e0fa20418c186628f9b6565a116017f988f2d9ae058480fae910a4659bf2", false, true},
		{"valid signature", peerCtx, "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4", "a2e3cf2037699b3856c72af280ab8501878495dd81595128df23ba3de0e52fd9126c02b9262b871074f5a34495cd1a1c13cf3d27881ce9a8846463b7d30024c37861e0fa20418c186628f9b6565a116017f988f2d9ae058480fae910a4659bf2", true, false},
	}

	service, err := newTestService()
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := service.JoinNetwork(tc.ctx, &types.JoinNetworkRequest{
				StakeAddress: tc.stakeAdr,
				PublicKey:    tc.pubKey,
				Signature:    tc.signature,
			})
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.valid, res.Result)
			}
		})
	}
}

func TestBlockOperation(t *testing.T) {
	service, err := newTestService()
	require.NoError(t, err)

	privKey := "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	priv := new(bls.SecretKey)
	err = priv.Unmarshal(common.FromHex(privKey))
	require.NoError(t, err)

	// join network
	peerCtx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: &net.IPAddr{},
	})

	res, err := service.JoinNetwork(peerCtx, &types.JoinNetworkRequest{
		StakeAddress: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		PublicKey:    "0x86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4",
		Signature:    "a2e3cf2037699b3856c72af280ab8501878495dd81595128df23ba3de0e52fd9126c02b9262b871074f5a34495cd1a1c13cf3d27881ce9a8846463b7d30024c37861e0fa20418c186628f9b6565a116017f988f2d9ae058480fae910a4659bf2",
	})

	require.NoError(t, err)
	require.True(t, res.Result)

	// get block
	_, err = service.GetBlock(context.Background(), &types.GetBlockRequest{BlockNumber: 0})
	require.Error(t, err)

	block, err := service.GetBlock(peerCtx, &types.GetBlockRequest{BlockNumber: 0})
	require.NoError(t, err)

	// commit block
	// wrong block number
	msg, err := proto.Marshal(block.Block)
	require.NoError(t, err)
	sig, err := priv.Sign(msg)
	require.NoError(t, err)
	sigMsg := sig.Serialize()
	_, err = service.CommitBlock(context.Background(), &types.CommitBlockRequest{BlockNumber: 0, Signature: common.Bytes2Hex(sigMsg[:])})
	require.Error(t, err)
	// last block number
	block, err = service.GetBlock(peerCtx, &types.GetBlockRequest{BlockNumber: 2})
	require.NoError(t, err)
	msg, err = proto.Marshal(block.Block)
	require.NoError(t, err)
	sig, err = priv.Sign(msg)
	require.NoError(t, err)
	sigMsg = sig.Serialize()
	_, err = service.CommitBlock(peerCtx, &types.CommitBlockRequest{BlockNumber: 2, Signature: common.Bytes2Hex(sigMsg[:])})
	require.NoError(t, err)
}
