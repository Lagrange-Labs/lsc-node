package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
	"log"
	"errors"
	json "encoding/json"

	host "github.com/libp2p/go-libp2p-core/host"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
//	accounts "github.com/ethereum/go-ethereum/accounts"
//	common "github.com/ethereum/go-ethereum/common"
)

var LOG_LEVEL int

// Placeholder - nodestaking contract address
const NODE_STAKING_ADDRESS = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
//const NODE_STAKING_ADDRESS = "0x00000000006c3852cbef3e08e8df289169ede581"	// test

// Placeholder - Return first Hardhat private key for now
var PRIVATE_KEY string = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
// Placeholder - Track staking listening to contract via rpc
var STAKE_STATE []string

type peerSummary struct {
	id		string
	address		string
	lastSeen	int64
}

var peerRegistry = map[string]peerSummary {}

// Returns private key as hex string.
func GetPrivateKey() string {
	return PRIVATE_KEY
}

// Sets private key in the form of a hex string.
func SetPrivateKey(privateKey string) {
	PRIVATE_KEY = privateKey
}

func main() {
	args := GetOpts()
	
	ks := args.keystore
	port := args.port
	stakingEndpoint := args.stakingEndpoint
	stakingWS := args.stakingWS
	attestEndpoint := args.attestEndpoint
	nick := args.nick
	peerAddr := args.peerAddr
	room := args.room
	leveldb := args.leveldb
	logLevel := args.logLevel
	_ = leveldb
	
	LOG_LEVEL = logLevel

	if(true) {} else
	if(ks == "") {
		privateKeyHex, publicKeyHex := GenerateKeypair()
		_ = publicKeyHex
		SetPrivateKey(privateKeyHex)
	} else {
		os.RemoveAll("./tmp/")
		store := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

		jsonBytes, err := ioutil.ReadFile(ks)
		if err != nil {
			log.Fatal(err)
		}

		input := Scan("Enter passphrase for keystore:")
		account, err := store.Import(jsonBytes, input, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(account)

		return
	}
	
	LogMessage(fmt.Sprintf("Port: %v",port),LOG_INFO)

	rpcStaking := LoadRpcClient(stakingEndpoint)
	ethStaking := LoadEthClient(stakingEndpoint)
	_ = rpcStaking
	_ = ethStaking

	rpcWS := LoadRpcClient(stakingWS)
	ethWS := LoadEthClient(stakingWS)
	_ = rpcWS
	_ = ethWS

	rpcAttest := LoadRpcClient(attestEndpoint)
	_ = rpcAttest
		
	ethAttestClients := LoadEthClientMulti(attestEndpoint)
	
	// Create listener
	node := CreateListener(port)

	if(len(nick) == 0) {
		nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), ShortID(node.ID()))
	}
	LogMessage("Nickname: "+nick,LOG_DEBUG)
	
	// Connect to Remote Peer
	ConnectRemote(node,peerAddr)
	
	// Core Routines
	go HeartBeat()

	// Messaging + Listening Routines	
	ps, topic, subscription := GetGossipSub(node,room)
	go HandleMessaging(node,topic,ps,nick,subscription)
	go ListenForBlocks(ethAttestClients,node,topic,ps,nick,subscription)
	// NodeStaking event listening
	go ListenForStaking(ethWS)

	// Sandbox - Contract Interaction
	CtrIntTest(rpcStaking,ethStaking)
//	go MineTest(rpcStaking)

	SendVerificationMessage(node,topic)

//	activeStakesTest(rpcStaking,ethStaking)
//	ethTest(eth)

        // SIGINT | SIGTERM Signal Handling - End
        TermHandler(node)
}

func HeartBeat() {
	for {
		t := GetUnixTimestamp()
		// Check peer registry every 10 seconds, drop if no updates in 30
		for id,peer := range peerRegistry {
			if(t - peer.lastSeen > 30) {
				LogMessage("Dropping peer '"+id+"' due to non-responsiveness",LOG_INFO)
				delete(peerRegistry,id)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func ProcessJoinMessage(message *GossipMessage) (error) {
	var jm = &JoinMessage{}
	err := json.Unmarshal([]byte(message.Data),jm)
	if(err != nil) {
		panic(err)
	}

	publicKey := jm.PublicKey
	genericMessage := jm.GenericMessage
	timestampStr := jm.Timestamp
	saltStr := jm.Salt
	sigtuple := jm.ECDSASignatureTuple
	_ = publicKey
	_ = genericMessage
	_ = timestampStr
	_ = saltStr
	_ = sigtuple
	
	signature, err := hexutil.Decode(sigtuple)
	if err != nil { panic(err) }
	_ = signature
	
	tuple := GenerateVerificationTupleFromJoinMessage(genericMessage, timestampStr, saltStr)
	_ = tuple
	
//	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	pubKey, err := crypto.SigToPub(KeccakHash(tuple), signature)
	if err != nil {
		panic(err)
	}
	_ = pubKey

	if common.HexToAddress(message.SenderNick) != crypto.PubkeyToAddress(*pubKey) {
		return errors.New("Failed to verify peer: "+crypto.PubkeyToAddress(*pubKey).String())
	}

	LogMessage("Peer verified: "+crypto.PubkeyToAddress(*pubKey).String(),LOG_NOTICE)
	return nil
}

const (
	MESSAGE_TYPE_JOINMESSAGE = 1
	MESSAGE_TYPE_STATEROOTMESSAGE = 2
)

func ProcessMessage(message *GossipMessage) (error) {
	var srm = &StateRootMessage{}
	switch message.Type {
		case "JoinMessage":
			return ProcessJoinMessage(message)
			break
		case "StateRootMessage":
			err := json.Unmarshal([]byte(message.Data),srm)
			fmt.Println(srm)
			_ = err
			return nil
			break
	}
	return errors.New("Invalid or unspecified message type.")
}

func TermHandler(node host.Host) {
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        <-ch
        LogMessage("Received signal, shutting down...",LOG_INFO)

        // shut the node down
        if err := node.Close(); err != nil {
                panic(err)
        }
        LogMessage("*DONE*",LOG_INFO)
}
