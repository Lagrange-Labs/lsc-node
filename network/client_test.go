package network

import (
	"context"
	"testing"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	require.Equal(t, addr.Hex(), "0x516D6C27C23CEd21BF7930E2a01F0BcA9A141a0d") // 0x516D6C27C23CEd21BF7930E2a01F0BcA9A141a0d
}

func TestBlockParams(t *testing.T) {
	cfg := &ClientConfig{
		EthereumURL:        "http://localhost:8545",
		CommitteeSCAddress: "0xF2740f6A6333c7B405aD7EfC68c74adAd83cC30D",
	}

	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		logger.Fatalf("failed to create the ethereum client: %v", err)
	}
	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
	if err != nil {
		logger.Fatalf("failed to create the committee smart contract: %v", err)
	}

	chainID, err := etherClient.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chain ID: %v", err)
	}

	params, err := committeeSC.CommitteeParams(nil, uint32(chainID.Uint64()))
	require.NoError(t, err)
	t.Log("params:", params)
}
