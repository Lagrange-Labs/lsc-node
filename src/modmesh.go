package main

import "fmt"
import "flag"
import "os"
import "os/signal"
import "syscall"
import "context"

import host "github.com/libp2p/go-libp2p-core/host"
import ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"

// const NODE_STAKING_ADDRESS = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

NODE_STAKING_ADDRESS := os.Getenv("NODE_STAKING_ADDRESS")

// Placeholder - Return first Hardhat private key for now
func getPrivateKey() string {
	// return "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	return os.Getenv("HARDHAT_PRIVATE_KEY")

}

func main() {
	// Parse Port
	portPtr := flag.Int("port",8081,"Server listening port")
	// Parse Nickname
	nickPtr := flag.String("nick","","Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later.")
	// Parse Room
	roomPtr := flag.String("room","rinkeby","Room / Network")
	// Parse Remote Peer
	peerAddrPtr := flag.String("peerAddr","","Remote Peer Address")
	// Parse ETH (Staking) URL
	stakingEndpointPtr := flag.String("stakingEndpoint","https://34.229.73.193:8545","Staking Endpoint URL:Port")
	// Parse ETH (Attestation) URL
	attestEndpointPtr := flag.String("attestEndpoint","https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79","Attestation Endpoint URL:Port")

	flag.Parse()

	port := *portPtr
	nick := *nickPtr
	room := *roomPtr
	peerAddr := *peerAddrPtr
	stakingEndpoint := *stakingEndpointPtr
	attestEndpoint := *attestEndpointPtr
	
	fmt.Println("Port:",port)

	rpcStaking := loadRpcClient(stakingEndpoint)
	ethStaking := loadEthClient(stakingEndpoint)

	rpcAttest := loadRpcClient(attestEndpoint)
	_ = rpcAttest
	ethAttest := loadEthClient(attestEndpoint)
	
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
	
	// Sandbox - Contract Interaction
	ctrIntTest(rpcStaking,ethStaking)
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
