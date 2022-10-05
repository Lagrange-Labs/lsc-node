package main

import "fmt"
import "os"
import "time"
import bufio "bufio"
import json "encoding/json"

import context "context"

import host "github.com/libp2p/go-libp2p-core/host"
import pubsub "github.com/libp2p/go-libp2p-pubsub"

// handleEvents runs an event loop that sends user input to the chat room
// and displays messages received from the chat room. It also periodically
// refreshes the list of peers in the UI.
func handleMessaging(node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription) {
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
		go readMessages(node,topic,ps,subscription);
			
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
	Message    string
	// peer controls a private key
	SenderID   string
	SenderNick string
}

const bufferSize = 4096

type MsgParams struct {
	ps *pubsub.PubSub
	topic *pubsub.Topic
	subscription *pubsub.Subscription
	node host.Host
	nick string
	message string
}

func readMessages(node host.Host, topic *pubsub.Topic, subscription *pubsub.PubSub, room *pubsub.Subscription) {
	messages :=  make(chan *GossipMessage, bufferSize)
	
	for {
		msg, err := room.Next(context.Background())
		if err != nil {
			close(messages)
			return
		}
		// only forward messages delivered by others
		if msg.ReceivedFrom == node.ID() {
			continue
		}
		fmt.Println(string(msg.Data))
		cm := new(GossipMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		messages <- cm
	}
}

func writeMessages(node host.Host, topic *pubsub.Topic, nick string, message string) error {
	m := GossipMessage{
		Message:    message,
		SenderID:   node.ID().Pretty(),
		SenderNick: nick,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(msgBytes))
	return topic.Publish(context.Background(), msgBytes)
}
