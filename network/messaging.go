package network

import (
	bufio "bufio"
	"bytes"
	"compress/gzip"
	json "encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	context "context"

	host "github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

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

// Converted to/from JSON and sent in the body of pubsub messages.
type GossipMessage struct {
	Type       string
	Data       string
	SenderID   string
	SenderNick string
}

const BufferSize = 4096

type MsgParams struct {
	ps           *pubsub.PubSub
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	node         host.Host
	nick         string
	message      string
}

func (lnode *LagrangeNode) ReadMessages() {
	node := lnode.node
	subscription := lnode.subscription

	messages := make(chan *GossipMessage, BufferSize)

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
		decompressedMessage, decompressErr := DecompressMessage(msg.Data)
		// LogMessage(fmt.Sprintf("ReadMessages: %v", string(decompressedMessage)), LOG_DEBUG)
		if decompressErr != nil {
			close(messages)
			return
		}

		// Parse message
		cm := new(GossipMessage)
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
			LogMessage(fmt.Sprintf("ReadMessages: %v", string(decompressedMessage)), LOG_DEBUG)

			processMessageError := lnode.ProcessMessage(cm)
			if processMessageError != nil {
				LogMessage(fmt.Sprintf("%v", processMessageError), LOG_WARNING)
				return
			}

			UpdatePeerRegistry(cm, msg)

			// send valid messages onto the Messages channel
			messages <- cm
		}
	}
}

func UpdatePeerRegistry(cm *GossipMessage, msg *pubsub.Message) {
	// Address
	fromId := fmt.Sprintf("%v", msg.ReceivedFrom)
	localTime := GetUnixTimestamp()
	summary := peerSummary{fromId, cm.SenderNick, localTime}
	// TODO use address instead of fromId once authentication and verification is ready
	// Counterpoint: additional processing required to derive address from public key.  Continue using fromID for time being.
	peerRegistry[cm.SenderNick] = summary
}

func WriteMessages(node host.Host, topic *pubsub.Topic, nick string, message string, messagetype string) error {
	m := GossipMessage{
		Type:       messagetype,
		Data:       message,
		SenderID:   node.ID().Pretty(),
		SenderNick: nick,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	compressedMessageBytes := CompressMessage(msgBytes)
	// LogMessage("WriteMessages: "+string(msgBytes), LOG_DEBUG)
	// LogMessage("I am here WriteMessages: "+string(compressedMessageBytes), LOG_DEBUG)
	return topic.Publish(context.Background(), compressedMessageBytes)
}

func CompressMessage(message []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(message); err != nil {
		return nil
	}
	if err := gz.Flush(); err != nil {
		return nil
	}
	if err := gz.Close(); err != nil {
		return nil
	}
	return b.Bytes()
}

func DecompressMessage(compressedMessage []byte) ([]byte, error) {
	b := bytes.NewReader(compressedMessage)
	gz, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, gz); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
