package consensus

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/store/memdb"
	"github.com/Lagrange-Labs/lagrange-node/testutil"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
)

func createTestState(t *testing.T) (*State, chan *sequencertypes.BlsSignature) {
	keystorePassword := "password"
	keystorePath := filepath.Join(t.TempDir(), "bls.json")
	err := testutil.GenerateRandomKeystore(string(crypto.BN254), keystorePassword, keystorePath)
	require.NoError(t, err)
	cfg := &Config{
		ProposerBLSKeystorePath:     keystorePath,
		ProposerBLSKeystorePassword: keystorePassword,
		RoundLimit:                  utils.TimeDuration(5 * time.Second),
		RoundInterval:               utils.TimeDuration(2 * time.Second),
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

func TestStateAggSignatur(t *testing.T) {
	rawBlock := "{\"_id\":\"649e0220951f3067cfef4a38\",\"agg_signature\":\"a6d26317d912da8327c60192d731ed9b33309dcae5f393f58caac1c9d4deaaaaa739e127de393984eafe6a353c42fd340ca3733c3b30af84725ceddf1bed7b11f3956274d9950857497c801ea7f491e3255bcc0480faf099ee8d8daced0e61b6\",\"block_header\":{\"current_committee\":\"2e3d2e5c97ee5320cccfd50434daeab6b0072558b693bb0e7f2eeca97741e514\",\"epoch_block_number\":9263248,\"next_committee\":\"2e3d2e5c97ee5320cccfd50434daeab6b0072558b693bb0e7f2eeca97741e514\",\"proposer_pub_key\":\"86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4\",\"proposer_signature\":\"b37d35c572578b68ef028699807c17e55f3eb49c9dd0a8bcd79cd2e32b5d2f3eda97c22dc0358f66de5401f4b4dae6ea00af4e143ee276b7c6d98ec57ead72d30d3e94e443d79b2537c431c58b097f5a70b63007b7eabe1aebfe0624a5ab7d61\",\"total_voting_power\":900000000},\"chain_header\":{\"block_hash\":\"0x3cf0f349f278d1ef6b54d07a463c675264d69202d5d409a97b307de244d6dedc\",\"block_number\":28809914,\"chain_id\":421613},\"pub_keys\":[\"86b50179774296419b7e8375118823ddb06940d9a28ea045ab418c7ecbe6da84d416cb55406eec6393db97ac26e38bd4\",\"b5695acd75a5d52e82eddf4ae1c01a1d456085da4ce255169cdac877b5a622386e7db6cb9c4c39b4eaf660dfa3a80d5d\",\"9496160c06e86aae28d251e082173b49ff10225f29611fae766d1a1a473d4bcf1188738208df163c01927b8c5df5160a\",\"91f8a57e971b831b80f1f6c04eed120c2c3ecb9d13adf7601e05614909158c3882bdbc04251cb689ecf653c28b981774\",\"90aa5215784d04f54d2df41cca557514c25b26cd0803da59c08eb3954633c4fe5131e195b1483123f5c9fe9f08200a6d\",\"837ca7f100239253d16d982102e47387a369cd1ba4bf9cf08ab85bd70c2b7559e7158196d9c60d03a571dfb3580b2b8e\",\"a771bcc2d948fa9801916e0c18aa27732f77d43b751138a3e119928103826427a5c11080620f9a9e37d8f4de3a4868cf\"]}"

	block := &sequencertypes.Block{}
	require.NoError(t, json.Unmarshal([]byte(rawBlock), block))

	pubkeys := make([]*bls.PublicKey, 0)
	for _, pk := range block.PubKeys {
		pubkey, err := utils.HexToBlsPubKey(pk)
		require.NoError(t, err)
		pubkeys = append(pubkeys, pubkey)
	}

	aggSignature, err := utils.HexToBlsSignature(block.AggSignature)
	require.NoError(t, err)
	msgHash := block.BlsSignature().Hash()
	verified, err := aggSignature.FastAggregateVerify(pubkeys, msgHash)
	require.NoError(t, err)
	require.True(t, verified)
}
