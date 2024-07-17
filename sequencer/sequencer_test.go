package sequencer

import (
	"os"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/optimism"
	"github.com/Lagrange-Labs/lagrange-node/store"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func TestFetchCommittee(t *testing.T) {
	ethURL := os.Getenv("ETH_MAIN_RPC")
	if len(ethURL) == 0 {
		t.Skip("ETH_RPC not set")
	}
	optURL := os.Getenv("OPT_MAIN_RPC")
	if len(optURL) == 0 {
		t.Skip("OPT_RPC not set")
	}

	seqConfig := Config{
		EthereumURL:          ethURL,
		Chain:                "optimism",
		CommitteeSCAddress:   "0xECc22f3EcD0EFC8aD77A78ad9469eFbc44E746F5",
		EigenDMSCAddress:     "0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A",
		StakingCheckInterval: utils.TimeDuration(5),
	}
	rpcConfig := rpcclient.Config{
		Optimism: &optimism.Config{
			RPCURLs:     []string{optURL},
			L1RPCURLs:   []string{ethURL},
			BeaconURL:   "http://localhost:8545",
			BatchInbox:  "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f",
			BatchSender: "0x0AEd0dC7f53CB452A34A3Fe4d6a7E4Fdd110ed0f",
		},
	}

	memDB, err := store.NewStorage(&store.Config{BackendType: "memdb"})
	require.NoError(t, err)
	s, err := NewSequencer(&seqConfig, &rpcConfig, memDB)
	require.NoError(t, err)

	c, err := s.fetchCommitteeRoot(75)
	require.NoError(t, err)
	require.NotNil(t, c)
	require.NoError(t, c.Verify())
}
