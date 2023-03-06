package node

import (
	"bufio"
	context "context"
	"crypto/ecdsa"
	json "encoding/json"
	"errors"
	"fmt"
	log "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"io/ioutil"
	//"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	host "github.com/libp2p/go-libp2p/core/host"

	"github.com/Lagrange-Labs/Lagrange-Node/bcclients"
	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/node/nodestaking"
	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	accounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
)

type peerSummary struct {
	id       string
	address  string
	lastSeen int64
}

var peerRegistry = map[string]peerSummary{}

func HeartBeat() {
	for {
		t := utils.GetUnixTimestamp()
		// Check peer registry every 10 seconds, drop if no updates in 30
		for id, peer := range peerRegistry {
			if t-peer.lastSeen > 30 {
				utils.LogMessage("Dropping peer '"+id+"' due to non-responsiveness", utils.LOG_INFO)
				delete(peerRegistry, id)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func UpdatePeerRegistry(cm *network.GossipMessage, msg *pubsub.Message) {
	// Address
	fromId := fmt.Sprintf("%v", msg.ReceivedFrom)
	localTime := utils.GetUnixTimestamp()
	summary := peerSummary{fromId, cm.SenderNick, localTime}
	// TODO use address instead of fromId once authentication and verification is ready
	// Counterpoint: additional processing required to derive address from public key.  Continue using fromID for time being.
	peerRegistry[cm.SenderNick] = summary
}

// Self-contained struct for managing the lagrange node and its configuration.
type LagrangeNode struct {
	// Derived objects
	node         host.Host
	topic        *pubsub.Topic
	pubsub       *pubsub.PubSub
	nick         string
	subscription *pubsub.Subscription

	rpcWS *rpc.Client
	ethWS *ethClient.Client

	rpcStaking          *rpc.Client
	ethStaking          *ethClient.Client
	nodeStakingInstance *nodestaking.Nodestaking

	rpcAttest *rpc.Client

	ethAttestClients []*ethClient.Client

	privateKey   string
	account      accounts.Account
	keystore     *keystore.KeyStore
	publicKeyHex string

	walletPath string
	address    common.Address
}

func NewLagrangeNode() *LagrangeNode {
	lnode := &LagrangeNode{}
	return lnode
}

func (lnode *LagrangeNode) Start(cfg *Config) {
	lnode.keystore.Unlock(lnode.account, "")

	peerAddr := cfg.PeerAddr
	room := cfg.Room
	leveldb := cfg.LevelDBPath
	_ = leveldb

	utils.LogMessage(fmt.Sprintf("Port: %v", cfg.Port), utils.LOG_INFO)

	lnode.rpcStaking = bcclients.LoadRpcClient(cfg.StakingEndpoint)
	lnode.ethStaking = bcclients.LoadEthClient(cfg.StakingEndpoint)
	lnode.nodeStakingInstance = GetStakingContract(lnode.ethStaking)

	lnode.rpcWS = bcclients.LoadRpcClient(cfg.StakingWS)
	lnode.ethWS = bcclients.LoadEthClient(cfg.StakingWS)
	lnode.rpcAttest = bcclients.LoadRpcClient(cfg.AttestEndpoint)
	lnode.ethAttestClients = bcclients.LoadEthClientMulti(cfg.AttestEndpoint)

	lnode.address = common.HexToAddress(cfg.StakerAddress)

	// Create listener
	node := network.CreateListener(cfg.Port)

	if len(cfg.Nickname) == 0 {
		lnode.nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), utils.ShortID(node.ID()))
	} else {
		lnode.nick = cfg.Nickname
	}
	utils.LogMessage("Nickname: "+cfg.Nickname, utils.LOG_DEBUG)

	// Connect to Remote Peer
	network.ConnectRemote(node, peerAddr)

	// Core Routines
	go HeartBeat()

	// Messaging + Listening Routines
	ps, topic, subscription := network.GetGossipSub(node, room)

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
	lnode.SimulateStaking(lnode.rpcStaking, lnode.ethStaking)
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

func (lnode *LagrangeNode) GetAddressString() string {
	return hexutil.Encode(lnode.account.Address.Bytes())
}

func (lnode *LagrangeNode) GetWalletPath() string {
	return lnode.walletPath
}

func (lnode *LagrangeNode) SetWalletPath(walletPath string) {
	lnode.walletPath = walletPath
}

func (lnode *LagrangeNode) HasAccount(address string) bool {
	return lnode.keystore.HasAddress(common.HexToAddress(address))
}

func (lnode *LagrangeNode) SetAddress(address string) {
	lnode.address = common.HexToAddress(address)
}

// handleEvents runs an event loop that sends user input to the chat room
// and displays messages received from the chat room. It also periodically
// refreshes the list of peers in the UI.
func (lnode *LagrangeNode) HandleMessaging() {
	/*
		peerRefreshTicker := time.NewTicker(time.Second)
		defer peerRefreshTicker.Stop()
	*/
	//	messages := make(chan *GossipMessage, bufferSize)
	reader := bufio.NewReader(os.Stdin)
	_ = reader
	for {
		//input, _ := reader.ReadString('\n')
		//go writeMessages(node,topic,nick,input)
		go lnode.ReadMessages()

		/*
			select {
			case m := <- messages:
				// when we receive a message from the chat room, print it to the message window
			case <-peerRefreshTicker.C:
				// refresh the list of peers in the chat room periodically
				ui.refreshPeers()
			case <- context.Background().Done():
				return
			case <-ui.doneCh:
				return
			}
		*/

		time.Sleep(1 * time.Second)
	}
}

func (lnode *LagrangeNode) ReadMessages() {
	node := lnode.node
	subscription := lnode.subscription

	messages := make(chan *network.GossipMessage, network.BufferSize)

	for {
		msg, err := subscription.Next(context.Background())
		if err != nil {
			close(messages)
			return
		}
		// Only forward messages delivered by others
		if msg.ReceivedFrom == node.ID() {
			continue
		}

		// Decompress message data
		decompressedMessage, decompressErr := network.DecompressMessage(msg.Data)
		// LogMessage(fmt.Sprintf("ReadMessages: %v", string(decompressedMessage)), LOG_DEBUG)
		if decompressErr != nil {
			close(messages)
			return
		}

		// Parse message
		cm := new(network.GossipMessage)
		err = json.Unmarshal(decompressedMessage, cm)
		if err != nil {
			continue
		}

		// 1. Verify peer.  Check if peer is authenticated, otherwise authenticate, otherwise ignore.
		val, ok := peerRegistry[cm.SenderNick]
		_ = val
		if !ok && cm.Type != "JoinMessage" {
			// Ignore
			continue
		} else {
			// 2. Process message (e.g., attestation).
			utils.LogMessage(fmt.Sprintf("ReadMessages: %v", string(decompressedMessage)), utils.LOG_DEBUG)

			processMessageError := lnode.ProcessMessage(cm)
			if processMessageError != nil {
				utils.LogMessage(fmt.Sprintf("%v", processMessageError), utils.LOG_WARNING)
				return
			}

			UpdatePeerRegistry(cm, msg)

			// send valid messages onto the Messages channel
			messages <- cm
		}
	}
}

func (lnode *LagrangeNode) ProcessJoinMessage(message *network.GossipMessage) error {
	var jm = &network.JoinMessage{}
	err := json.Unmarshal([]byte(message.Data), jm)
	if err != nil {
		panic(err)
	}

	genericMessage := jm.GenericMessage
	timestampStr := jm.Timestamp
	saltStr := jm.Salt
	sigtuple := jm.ECDSASignatureTuple
	_ = genericMessage
	_ = timestampStr
	_ = saltStr
	_ = sigtuple

	signature, err := hexutil.Decode(sigtuple)
	if err != nil {
		panic(err)
	}
	_ = signature

	tuple := network.GenerateVerificationTupleFromJoinMessage(genericMessage, timestampStr, saltStr)
	_ = tuple

	//	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	pubKey, err := crypto.SigToPub(bcclients.KeccakHash(tuple), signature)
	if err != nil {
		panic(err)
	}
	_ = pubKey

	signatureVerified := common.HexToAddress(message.SenderNick) == crypto.PubkeyToAddress(*pubKey)
	addr := crypto.PubkeyToAddress(*pubKey)
	if !signatureVerified {
		return errors.New("Failed to verify signature for peer: " + addr.String())
	}

	stakeVerified := lnode.VerifyStake(addr)
	fmt.Println(stakeVerified)
	if !stakeVerified {
		return errors.New("Failed to verify stake for peer: " + addr.String())
	}

	utils.LogMessage("Peer verified: "+crypto.PubkeyToAddress(*pubKey).String(), utils.LOG_NOTICE)
	return nil
}

const (
	MESSAGE_TYPE_JOINMESSAGE      = 1
	MESSAGE_TYPE_STATEROOTMESSAGE = 2
)

func (lnode *LagrangeNode) ProcessMessage(message *network.GossipMessage) error {
	var srm = &StateRootMessage{}
	switch message.Type {
	case "JoinMessage":
		return lnode.ProcessJoinMessage(message)
		break
	case "StateRootMessage":
		err := json.Unmarshal([]byte(message.Data), srm)
		fmt.Println(srm)
		_ = err
		return nil
		break
	}
	return errors.New("Invalid or unspecified message type.")
}

func (lnode *LagrangeNode) TermHandler() {

	lnode.SimulateUnstaking(lnode.rpcStaking, lnode.ethStaking)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	utils.LogMessage("Received signal, shutting down...", utils.LOG_INFO)

	// shut the node down
	if err := lnode.node.Close(); err != nil {
		panic(err)
	}
	utils.LogMessage("*DONE*", utils.LOG_INFO)
}

func (lnode *LagrangeNode) SendVerificationMessage() {
	node := lnode.node
	topic := lnode.topic

	// ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt

	tuple, timestampStr, genericMessage, saltStr := network.GenerateVerificationTuple()

	tupleHash := bcclients.KeccakHash(tuple)

	signatureTuple, err := lnode.keystore.SignHash(lnode.account, tupleHash)
	if err != nil {
		panic(err)
	}

	signatureHex := hexutil.Encode(signatureTuple)

	joinMessage := network.JoinMessage{
		GenericMessage:      genericMessage,
		Timestamp:           timestampStr,
		Salt:                saltStr,
		ECDSASignatureTuple: signatureHex}

	json, err := json.Marshal(joinMessage)
	if err != nil {
		panic(err)
	}
	bytes := []byte(json)
	msg := string(bytes)

	network.WriteMessages(node, topic, lnode.GetAddressString(), msg, "JoinMessage")
}

func (lnode *LagrangeNode) LoadKeystore() (accounts.Account, *keystore.KeyStore) {
	ks := keystore.NewKeyStore(lnode.GetWalletPath(), keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Find(accounts.Account{Address: lnode.address})
	if err != nil {
		utils.LogMessage(fmt.Sprintf("%v", err), utils.LOG_WARNING)
	} else {
		utils.LogMessage(fmt.Sprintf("Keystore loaded for address %v", account.Address), utils.LOG_NOTICE)
	}
	return account, ks
}

func (lnode *LagrangeNode) LoadAccount() {
	account, ks := lnode.LoadKeystore()
	lnode.account = account
	lnode.keystore = ks
}

// Generates and returns keystore from private key.
func (lnode *LagrangeNode) InitKeystore(privateKey *ecdsa.PrivateKey) (accounts.Account, *keystore.KeyStore) {
	ks := keystore.NewKeyStore(lnode.GetWalletPath(), keystore.StandardScryptN, keystore.StandardScryptP)
	// No password until key management strategy established
	//input := Scan("Enter passphrase for new keystore:")
	//account,err := ks.ImportECDSA(privateKey,input)
	account, err := ks.ImportECDSA(privateKey, "")
	if err != nil {
		panic(err)
	}
	utils.LogMessage(fmt.Sprintf("New keystore created for address %v", account.Address), utils.LOG_NOTICE)
	utils.LogMessage(fmt.Sprintf("URL: %v", account.URL), utils.LOG_NOTICE)
	return account, ks
}

func (lnode *LagrangeNode) GenerateAccount() {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	lnode.GenerateAccountFromPrivateKey(privateKey)
}

func (lnode *LagrangeNode) GenerateAccountFromPrivateKey(privateKey *ecdsa.PrivateKey) {
	account, keystore := lnode.InitKeystore(privateKey)
	lnode.account = account
	lnode.keystore = keystore
}

func (lnode *LagrangeNode) GenerateAccountFromPrivateKeyString(privateKeyString string) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		panic(err)
	}
	lnode.GenerateAccountFromPrivateKey(privateKey)
}

func (lnode *LagrangeNode) GetAuth() *bind.TransactOpts {
	auth, err := bind.NewKeyStoreTransactor(lnode.keystore, lnode.account)
	if err != nil {
		panic(err)
	}
	return auth
}
