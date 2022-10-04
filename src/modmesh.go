package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"context"
	"io/ioutil"
	"log"

	host "github.com/libp2p/go-libp2p-core/host"
	ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
//	accounts "github.com/ethereum/go-ethereum/accounts"
//	common "github.com/ethereum/go-ethereum/common"
)

const NODE_STAKING_ADDRESS = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
//const NODE_STAKING_ADDRESS = "0x00000000006c3852cbef3e08e8df289169ede581"	// test

// Placeholder - Return first Hardhat private key for now
var PRIVATE_KEY string = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
// Placeholder - Track staking listening to contract via rpc
var STAKE_STATE []string

func getPrivateKey() string {
	return PRIVATE_KEY
}

func setPrivateKey(privateKey string) {
	PRIVATE_KEY = privateKey
}

func main() {
	args := getOpts()
	
	ks := args.keystore
	port := args.port
	stakingEndpoint := args.stakingEndpoint
	stakingWS := args.stakingWS
	attestEndpoint := args.attestEndpoint
	nick := args.nick
	peerAddr := args.peerAddr
	room := args.room

	if(true) {} else
	if(ks == "") {
		privateKeyHex, publicKeyHex := generateKeypair()
		_ = publicKeyHex
		setPrivateKey(privateKeyHex)
	} else {
		os.RemoveAll("./tmp/")
		store := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

		jsonBytes, err := ioutil.ReadFile(ks)
		if err != nil {
			log.Fatal(err)
		}

		input := scan("Enter passphrase for keystore:")
		account, err := store.Import(jsonBytes, input, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(account)

		return
	}
	
	fmt.Println("Port:",port)

	rpcStaking := loadRpcClient(stakingEndpoint)
	ethStaking := loadEthClient(stakingEndpoint)
	_ = rpcStaking
	_ = ethStaking

	rpcWS := loadRpcClient(stakingWS)
	ethWS := loadEthClient(stakingWS)
	_ = rpcWS
	_ = ethWS

	rpcAttest := loadRpcClient(attestEndpoint)
	_ = rpcAttest
	ethAttest := loadEthClient(attestEndpoint)
	_ = ethAttest
	// Create listener
	node := createListener(port)

	if(len(nick) == 0) {
		nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(node.ID()))
	}
	fmt.Println("Nickname:",nick)

	// Get P2P Address Info
	localInfo := getAddrInfo(node);
	_ = localInfo

	// Ping test - please determine an approach to finding peers, rather than self-pinging	
	ch := ping.Ping(context.Background(), node, localInfo.ID)
	for i := 0; i < 5; i++ {
		res := <-ch
		fmt.Println("Got ping response.", "Latency:", res.RTT)
	}
	
	// Connect to Remote Peer
	connectRemote(node,peerAddr)
	
	ps, topic, subscription := getGossipSub(node,room)
	go handleMessaging(node,topic,ps,nick,subscription)
	go listenForBlocks(ethAttest,node,topic,ps,nick,subscription)
	go mineTest(rpcStaking)
//	os.Exit(0)
	
	// Sandbox - Contract Interaction
//	ctrIntTest(rpcStaking,ethStaking)

//	go listenForStaking(ethWS)

//	activeStakesTest(rpcStaking,ethStaking)
//	ethTest(eth)

        // SIGINT | SIGTERM Signal Handling - End
        termHandler(node)
}

func termHandler(node host.Host) {
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        <-ch
        fmt.Println("Received signal, shutting down...")

        // shut the node down
        if err := node.Close(); err != nil {
                panic(err)
        }
}
