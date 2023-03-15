package node

import (
	context "context"
	"fmt"
	log "log"
	"math/big"

	"github.com/Lagrange-Labs/Lagrange-Node/bcclients"
	"github.com/Lagrange-Labs/Lagrange-Node/node/nodestaking"
	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
)

// Placeholder - nodestaking contract address
const NODE_STAKING_ADDRESS = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

//const NODE_STAKING_ADDRESS = "0x00000000006c3852cbef3e08e8df289169ede581"	// test

const (
	STAKE_STATUS_CLOSED  = 0
	STAKE_STATUS_PENDING = 1
	STAKE_STATUS_OPEN    = 2
)

type stakeSummary struct {
	address string
	chainId *big.Int
}

var stakeRegistry = map[string]stakeSummary{}

// Returns a NodeStaking struct instance.
func GetStakingContract(client *ethClient.Client) *nodestaking.Nodestaking {
	address := common.HexToAddress(NODE_STAKING_ADDRESS)
	instance, err := nodestaking.NewNodestaking(address, client)
	if err != nil {
		log.Fatal(err)
	}
	utils.LogMessage("Loaded contract address "+fmt.Sprintf("%v", address), utils.LOG_DEBUG)
	return instance
}

// Handler function for NodeStaking smart contract events detected while listening to network.
func HandleStakingEvent(vLog *nodestaking.NodestakingStakedNode) {
	node := vLog.Node
	chainId := vLog.ChainId
	amount := vLog.Amount
	claimTime := vLog.ClaimTime

	utils.LogMessage(fmt.Sprintf("Staking Event: node=%v, chainId=%v, amount=%v, claimTime=%v", node, chainId, amount, claimTime), utils.LOG_NOTICE)
	address := node.Hex()
	stakeSummary := stakeSummary{
		address: address,
		chainId: chainId}
	stakeRegistry[address] = stakeSummary
}

// Listens to network for NodeStaking smart contract events and handles accordingly.
func (lnode *LagrangeNode) ListenForStaking() {
	ethWS := lnode.ethWS
	// For local hardhat testing, use --stakingWS="ws://0.0.0.0:8545"
	sc := GetStakingContract(ethWS)
	_ = sc
	logs := make(chan *nodestaking.NodestakingStakedNode)

	block := uint64(100)

	sub, err := sc.WatchStakedNode(&bind.WatchOpts{&block, context.Background()}, logs)
	if err != nil {
		panic(err)
	}

	utils.LogMessage("Listening to contract "+NODE_STAKING_ADDRESS, utils.LOG_INFO)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
			continue
		case vLog := <-logs:
			HandleStakingEvent(vLog)
		}
	}
}

// Initializes staking transaction with NodeStaking smart contract.
func StakeBegin(instance *nodestaking.Nodestaking) *big.Int {
	stake, err := instance.StakeAmount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stake amount:", stake)
	return stake
}

// Adds stake to NodeStaking smart contract.
// Staking transaction must first be initialized with StakeBegin().
func StakeAdd(instance *nodestaking.Nodestaking, auth *bind.TransactOpts) {
	res, err := instance.AddStake(auth, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	_ = res
}

// Begins removal of stake from NodeStaking smart contract.
func StakeRemoveBegin(instance *nodestaking.Nodestaking, auth *bind.TransactOpts) {
	res, err := instance.StartStakeRemoval(auth, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	_ = res
}

// Finalizes removal of stake from NodeStaking smart contract.
// Staking removal transaction must first be initialzied with StakeRemoveBegin().
func StakeRemoveFinish(instance *nodestaking.Nodestaking, auth *bind.TransactOpts) {
	res, err := instance.FinishStakeRemoval(auth, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	_ = res
}

// Returns true if address provided is currently staked; false if not staked, or in the process of unstaking.
func (lnode *LagrangeNode) VerifyStake(addr common.Address) bool {
	client := lnode.ethStaking
	instance := GetStakingContract(client)

	activeStakes, err := instance.ActiveStakes(&bind.CallOpts{}, addr, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	return activeStakes == STAKE_STATUS_OPEN
}

func ActiveStakesTest(rpc *rpc.Client, client *ethClient.Client) {
	instance := GetStakingContract(client)
	_ = instance
	// activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},common.Address(""),big.NewInt(4))
	// if(err != nil) { panic(err) }
	// fmt.Println(activestakes)
}

// Reference function walking through NodeStaking contract's staking and unstaking transactions
func (lnode *LagrangeNode) SimulateStaking(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := lnode.nodeStakingInstance

	// Retrieve private key, public key, address
	fromAddress := lnode.account.Address

	utils.LogMessage("Testing staking for address "+fromAddress.String(), utils.LOG_NOTICE)

	// Cleanup previous session if necessary
	if lnode.VerifyStake(fromAddress) {
		lnode.SimulateUnstaking(rpc, client)
	}

	// Verify Stake
	isStaked := lnode.VerifyStake(fromAddress)
	fmt.Println("Stake Verification:", isStaked)

	// Request nonce for transaction
	nonce := bcclients.GetNonce(client, fromAddress)

	// Request gas price
	gasPrice := bcclients.GetGasPrice(client)

	// Begin Staking Transaction
	stake := StakeBegin(instance)

	auth := lnode.GetAuth()
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = stake
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Verify Stake
	isStaked = lnode.VerifyStake(fromAddress)
	fmt.Println("Stake Verification:", isStaked)

	// Add Stake
	StakeAdd(instance, auth)

	// Verify Stake
	isStaked = lnode.VerifyStake(fromAddress)
	fmt.Println("Stake Verification:", isStaked)

	// Hardhat - Mine Blocks
	bcclients.MineBlocks(rpc, 5)
}

func (lnode *LagrangeNode) SimulateUnstaking(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := lnode.nodeStakingInstance

	// Retrieve private key, public key, address
	fromAddress := lnode.account.Address

	// Request gas price
	gasPrice := bcclients.GetGasPrice(client)

	auth := lnode.GetAuth()
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Update Nonce and Val
	nonce := bcclients.GetNonce(client, fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	// Begin Stake Removal
	StakeRemoveBegin(instance, auth)

	// Verify Stake
	isStaked := lnode.VerifyStake(fromAddress)
	fmt.Println("Stake Verification:", isStaked)

	// Hardhat - Mine More Blocks
	bcclients.MineBlocks(rpc, 5)

	// Update Nonce
	nonce = bcclients.GetNonce(client, fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))

	// Finalize Stake Removal
	StakeRemoveFinish(instance, auth)

	// Verify Stake
	isStaked = lnode.VerifyStake(fromAddress)
	fmt.Println("Stake Verification:", isStaked)

	fmt.Println("End staking test.")
}
