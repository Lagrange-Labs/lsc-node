package network

import (
	"fmt"
	"testing"
	"time"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/goleveldb"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestECDSASignVerify(t *testing.T) {
	// sign the BlsSignature
	privateKey, err := crypto.HexToECDSA("232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429")
	require.NoError(t, err)

	signature := &sequencertypes.BlsSignature{
		ChainHeader: &sequencertypes.ChainHeader{
			ChainId:     1337,
			BlockHash:   "0xafe58890693444d9116c940a5ff4418723e7f75869b30c9d8e4528e147cb4b7f",
			BlockNumber: 3,
		},
		CurrentCommittee: "0x9c11dac30afc6d443066d31976ece1015527da8d1c6f5e540ce649970f2e9129",
		NextCommittee:    "0x0538f196c8c36715f077e40f62b62795d83a4d82fddff30511375c9f6917a26b",
		BlsSignature:     "b3ad75be8554f25871e395268a2aec2d1d65003e70d4cd5b1560f37a85c7917fb82d66e22829c333043b4d6c3434151b13fb6b60d06f150132390f177c7891e97213c34cc843937f5e372035dcbb8be32ba6bf61a1545bdc2aafabd0fb60c5a4",
	}

	reqMsg := contypes.GetCommitRequestHash(signature)
	sig, err := crypto.Sign(reqMsg, privateKey)
	require.NoError(t, err)
	t.Log("signature:", common.Bytes2Hex(sig))
	// verify the signature
	isVerified, addr, err := utils.VerifyECDSASignature(reqMsg, sig)
	require.NoError(t, err)
	require.True(t, isVerified)
	require.Equal(t, addr.Hex(), "0x516D6C27C23CEd21BF7930E2a01F0BcA9A141a0d")
}

var _ rpctypes.RpcClient = (*mockRPC)(nil)

type mockRPC struct {
	chBatch            chan *sequencerv2types.BatchHeader
	chBeginBlockNumber chan uint64
}

func (m *mockRPC) GetCurrentBlockNumber() (uint64, error) {
	return 0, nil
}

func (m *mockRPC) GetFinalizedBlockNumber() (uint64, error) {
	return 0, nil
}

func (m *mockRPC) GetChainID() (uint32, error) {
	return 0, nil
}

func (m *mockRPC) SetBeginBlockNumber(l1BlockNumber uint64) bool {
	m.chBeginBlockNumber <- l1BlockNumber
	return true
}

func (m *mockRPC) NextBatch() (*sequencerv2types.BatchHeader, error) {
	batch, ok := <-m.chBatch
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}
	return batch, nil
}

func TestClientStorage(t *testing.T) {
	db, err := goleveldb.NewDB(t.TempDir())
	require.NoError(t, err)
	chBatch := make(chan *sequencerv2types.BatchHeader, 10)
	chBeginBlockNumber := make(chan uint64, 1)
	client := &Client{
		rpcClient: &mockRPC{
			chBatch:            chBatch,
			chBeginBlockNumber: chBeginBlockNumber,
		},
		db: db,
	}

	go func() {
		client.startBatchFetching()
	}()

	// push some batches
	for i := 1; i <= 10; i++ {
		chBatch <- &sequencerv2types.BatchHeader{
			L1BlockNumber: uint64(i),
			L1TxIndex:     1,
			L2Blocks: []*sequencerv2types.BlockHeader{
				{
					BlockNumber: uint64(i),
				},
			},
		}
		time.Sleep(200 * time.Millisecond)
	}

	// get previous batch
	prev, err := client.getPrevBatchL1Number(3, 0)
	require.NoError(t, err)
	require.Equal(t, uint64(2), prev)
	prev, err = client.getPrevBatchL1Number(3, 1)
	require.NoError(t, err)
	require.Equal(t, uint64(2), prev)
	prev, err = client.getPrevBatchL1Number(8, 2)
	require.NoError(t, err)
	require.Equal(t, uint64(8), prev)

	// get batch by L1 block number
	_, err = client.getBatchHeader(3, 2)
	require.Error(t, err)
	batch, err := client.getBatchHeader(5, 5)
	require.NoError(t, err)
	require.Equal(t, uint64(5), batch.L1BlockNumber)
	require.Equal(t, uint32(1), batch.L1TxIndex)

	// init begin block number
	err = client.initBeginBlockNumber(5)
	require.NoError(t, err)
	beginBlockNumber := <-chBeginBlockNumber
	require.Equal(t, uint64(10), beginBlockNumber)

	err = client.initBeginBlockNumber(11)
	require.NoError(t, err)
	beginBlockNumber = <-chBeginBlockNumber
	require.Equal(t, uint64(11), beginBlockNumber)
}
