package governance

import (
	"fmt"
	"github.com/Lagrange-Labs/lagrange-node/governance/nodestaking"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	trie "github.com/Lagrange-Labs/Lagrange-Node/trie"
)

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	stackingSC      *nodestaking.Nodestaking
	stakingInterval uint32
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *Config) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	stakingSC, err := nodestaking.NewNodestaking(common.HexToAddress(cfg.StakingSCAddress), client)
	if err != nil {
		return nil, err
	}
	return &Governance{
		stackingSC:      stakingSC,
		stakingInterval: cfg.StakingCheckInterval,
	}, nil
}

/*
// Reference implementation - update committees
func (gov *Governance) UpdateCommittees() string {
	instance := gov.stackingSC
	ethStaking, err := ethclient.Dial(cfg.EthereumURL)
	
	// Compute claim time for committee calculation time
	claimTime, err:= instance.ClaimDelay(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	// Determine staking block
	stakingBlock,err := ethStaking.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	
	// Update committee if applicable
	if int(stakingBlock.Number().Int64()) % int(claimTime.Int64()) == 0 {
		// "Current" Committee Root - Derive from cache
		//lnode.committeeRoot = lnode.nextCommitteeRoot
		//fmt.Println("Committee Root:",hexutil.Encode(lnode.committeeRoot))
		// "Next" Committee Root - Derive from contract
		//ncr := lnode.GetNextCommitteeRoot()
		//lnode.nextCommitteeRoot = ncr
		//fmt.Println("Next Committee Root:",hexutil.Encode(ncr))
	}
}
*/

func (gov *Governance) GetNextCommitteeRoot(cfg *Config) []byte {
	ethStaking := ethclient.Dial(cfg.EthereumURL)
	
	// Pull number of actively staked addresses
	instance := gov.stackingSC
	numStaked,err := instance.GetActiveStakeAddrsLength(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	numStakedInt := int(numStaked.Int64())
	var stakedAddrs []common.Address
	
	// Pull actively staked addresses
	fmt.Println("numStakedInt:",numStakedInt)
	
	chainID,err := ethStaking.NetworkID(context.Background())
	if err != nil { panic(err) }
	
	for i := 0; i < numStakedInt; i++ {
		// 1. Pull address at index i
		addr,err := instance.ActiveStakeAddrs(&bind.CallOpts{},big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		// 2. Verify that address corresponds to ChainID correspondin to network being attested to
		stakeStatus,err := instance.ActiveStakes(&bind.CallOpts{},addr,chainID)
		if err != nil { panic(err) }
		fmt.Println("stakestatus:",int(stakeStatus),"chainID:",chainID)
		if int(stakeStatus) != STAKE_STATUS_OPEN {
			continue
		}
		// 3. Append address list
		stakedAddrs = append(stakedAddrs,addr)
	}
	
	// Sort addresses
	sort.Slice(stakedAddrs, func(i int, j int) bool {
		return stakedAddrs[i].Hex() < stakedAddrs[j].Hex()
	})
	
	fmt.Println("stakedAddrs:",stakedAddrs)

	ptrie := trie.NewTrie()

	for _,addr := range stakedAddrs {
		ptrie.Add(addr.Bytes())
	}
	//lnode.nextCommitteeRoot = ptrie.Hash()
	return ptrie.Hash()
}
