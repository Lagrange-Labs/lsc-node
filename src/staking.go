package main

import (
	"fmt"
	"math/big"
	common "github.com/ethereum/go-ethereum/common"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
	log "log"
	context "context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	STAKE_STATUS_CLOSED = 0
	STAKE_STATUS_PENDING = 1
	STAKE_STATUS_OPEN = 2
)

type stakeSummary struct {
	address		string
	chainId		*big.Int
}

var stakeRegistry = map[string]stakeSummary {}

// Returns a NodeStaking struct instance.
func GetStakingContract(client *ethClient.Client) *Nodestaking {
	address := common.HexToAddress(NODE_STAKING_ADDRESS)
	instance, err := NewNodestaking(address,client)
	if err != nil {
		log.Fatal(err)
	}
	LogMessage("Loaded contract address "+fmt.Sprintf("%v",address),LOG_INFO)
	return instance
}

// Handler function for NodeStaking smart contract events detected while listening to network.
func HandleStakingEvent(vLog *NodestakingStakedNode) {
	node := vLog.Node
	chainId := vLog.ChainId
	amount := vLog.Amount
	claimTime := vLog.ClaimTime

	LogMessage(fmt.Sprintf("Staking Event: node=%v, chainId=%v, amount=%v, claimTime=%v",node,chainId,amount,claimTime),LOG_NOTICE)
	address := node.Hex()
	stakeSummary := stakeSummary {
		address: address,
		chainId: chainId }
	stakeRegistry[address] = stakeSummary
}

// Listens to network for NodeStaking smart contract events and handles accordingly.
func ListenForStaking(ethWS *ethClient.Client) {
	// For local hardhat testing, use --stakingWS="ws://0.0.0.0:8545"
	sc := GetStakingContract(ethWS)
	_ = sc
	logs := make(chan *NodestakingStakedNode)

	block := uint64(100)
	
	sub,err := sc.WatchStakedNode(&bind.WatchOpts{&block,context.Background()},logs)
	if err != nil { panic(err) }
	
	LogMessage("Listening to contract "+NODE_STAKING_ADDRESS,LOG_INFO)
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
func StakeBegin(instance *Nodestaking) *big.Int {
	stake, err := instance.StakeAmount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stake amount:",stake)
	return stake
}

// Adds stake to NodeStaking smart contract.
// Staking transaction must first be initialized with StakeBegin().
func StakeAdd(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.AddStake(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

// Begins removal of stake from NodeStaking smart contract.
func StakeRemoveBegin(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.StartStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

// Finalizes removal of stake from NodeStaking smart contract.
// Staking removal transaction must first be initialzied with StakeRemoveBegin().
func StakeRemoveFinish(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.FinishStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

// Returns true if address provided is currently staked; false if not staked, or in the process of unstaking.
func VerifyStake(client *ethClient.Client, instance *Nodestaking, addr common.Address) bool {
	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},addr,big.NewInt(4))
	if(err != nil) { panic(err) }
	return activeStakes == STAKE_STATUS_OPEN
}

func ActiveStakesTest(rpc *rpc.Client, client *ethClient.Client) {
	instance := GetStakingContract(client)
	_ = instance
//	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},common.Address(""),big.NewInt(4))
//	if(err != nil) { panic(err) }
//	fmt.Println(activestakes)
}

// Reference function walking through NodeStaking contract's staking and unstaking transactions
func CtrIntTest(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := GetStakingContract(client)

	// Retrieve private key, public key, address
	credentials := GetCredentials()
	privateKey := credentials.privateKeyECDSA
	fromAddress := credentials.address

	// Verify Stake
	isStaked := VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Request nonce for transaction	
	nonce := GetNonce(client,fromAddress)

	// Request gas price
	gasPrice := GetGasPrice(client)

	// Begin Staking Transaction
	stake := StakeBegin(instance)

	auth := GetAuth(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = stake
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Add Stake
	StakeAdd(instance,auth)
	
	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Hardhat - Mine Blocks
	MineBlocks(rpc,5)

	// Update Nonce and Val
	nonce = GetNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	
	// Begin Stake Removal
	StakeRemoveBegin(instance, auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Hardhat - Mine More Blocks
	MineBlocks(rpc,5)
	
	// Update Nonce
	nonce = GetNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	
	// Finalize Stake Removal
	StakeRemoveFinish(instance,auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)
	
	fmt.Println("End staking test.")
}
