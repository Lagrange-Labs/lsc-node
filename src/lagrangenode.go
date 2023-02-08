package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	//"log"

	host "github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"

	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
	accounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Self-contained struct for managing the lagrange node and its configuration.

type LagrangeNode struct {
	// Opts
	opts *Opts
	
	// Derived objects
	node host.Host
	topic *pubsub.Topic
	pubsub *pubsub.PubSub
	nick string
	subscription *pubsub.Subscription
	
	rpcWS *rpc.Client
	ethWS *ethClient.Client
	
	rpcStaking *rpc.Client
	ethStaking *ethClient.Client
	nodeStakingInstance *Nodestaking
	
	rpcAttest *rpc.Client
	
	ethAttestClients []*ethClient.Client
	
	privateKey string
	account accounts.Account
	keystore *keystore.KeyStore
	publicKeyHex string
}

func NewLagrangeNode() *LagrangeNode {
	lnode := &LagrangeNode{}
	return lnode
}

func (lnode *LagrangeNode) Start() {
	lnode.keystore.Unlock(lnode.account,"")

	peerAddr := lnode.opts.peerAddr
	room := lnode.opts.room
	leveldb := lnode.opts.leveldb
	logLevel := lnode.opts.logLevel
	_ = leveldb
	
	LOG_LEVEL = logLevel

	LogMessage(fmt.Sprintf("Port: %v",lnode.opts.port),LOG_INFO)

	lnode.rpcStaking = LoadRpcClient(lnode.opts.stakingEndpoint)
	lnode.ethStaking = LoadEthClient(lnode.opts.stakingEndpoint)
	lnode.nodeStakingInstance = GetStakingContract(lnode.ethStaking)
	
	lnode.rpcWS = LoadRpcClient(lnode.opts.stakingWS)
	lnode.ethWS = LoadEthClient(lnode.opts.stakingWS)
	lnode.rpcAttest = LoadRpcClient(lnode.opts.attestEndpoint)
	lnode.ethAttestClients = LoadEthClientMulti(lnode.opts.attestEndpoint)
	
	// Create listener
	node := CreateListener(lnode.opts.port)

	if(len(lnode.opts.nick) == 0) {
		lnode.nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), ShortID(node.ID()))
	} else {
		lnode.nick = lnode.opts.nick
	}
	LogMessage("Nickname: "+lnode.opts.nick,LOG_DEBUG)
	
	// Connect to Remote Peer
	ConnectRemote(node,peerAddr)
	
	// Core Routines
	go HeartBeat()

	// Messaging + Listening Routines	
	ps, topic, subscription := GetGossipSub(node,room)

	//node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription
	lnode.node = node
	lnode.topic = topic
	lnode.pubsub = ps
	lnode.subscription = subscription

	go lnode.HandleMessaging()
	go lnode.ListenForBlocks()
	// NodeStaking event listening
	go lnode.ListenForStaking()

	// Sandbox - Contract Interaction
	lnode.SimulateStaking(lnode.rpcStaking,lnode.ethStaking)
//	go MineTest(rpcStaking)

	lnode.SendVerificationMessage()

//	activeStakesTest(rpcStaking,ethStaking)
//	ethTest(eth)

        // SIGINT | SIGTERM Signal Handling - End
        lnode.TermHandler()
}

func (lnode *LagrangeNode) Stop() {
	lnode.keystore.Lock(lnode.account.Address)
	lnode.TermHandler()
}

func (lnode *LagrangeNode) SetOpts(opts *Opts) {
	lnode.opts = opts
}

func (lnode *LagrangeNode) GetAddressString() string {
	return hexutil.Encode(lnode.account.Address.Bytes())
}
