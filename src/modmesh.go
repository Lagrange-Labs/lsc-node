package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"errors"
	json "encoding/json"

	host "github.com/libp2p/go-libp2p-core/host"
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
	lnode := NewLagrangeNode()
	lnode.SetOpts(args)
	lnode.Start()	
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

func (lnode *LagrangeNode) ProcessJoinMessage(message *GossipMessage) (error) {
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
	
	signatureVerified := common.HexToAddress(message.SenderNick) == crypto.PubkeyToAddress(*pubKey)
	addr := crypto.PubkeyToAddress(*pubKey)
	if !signatureVerified {
		return errors.New("Failed to verify signature for peer: "+addr.String())
	}
	
	stakeVerified := lnode.VerifyStake(addr)
	fmt.Println(stakeVerified)
	if !stakeVerified {
		return errors.New("Failed to verify stake for peer: "+addr.String())		
	}

	LogMessage("Peer verified: "+crypto.PubkeyToAddress(*pubKey).String(),LOG_NOTICE)
	return nil
}

const (
	MESSAGE_TYPE_JOINMESSAGE = 1
	MESSAGE_TYPE_STATEROOTMESSAGE = 2
)

func (lnode *LagrangeNode) ProcessMessage(message *GossipMessage) (error) {
	var srm = &StateRootMessage{}
	switch message.Type {
		case "JoinMessage":
			return lnode.ProcessJoinMessage(message)
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

func (lnode *LagrangeNode) TermHandler(node host.Host) {

	lnode.SimulateUnstaking(lnode.rpcStaking,lnode.ethStaking)	

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
