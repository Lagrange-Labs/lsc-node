package network

import (
	"bytes"
	"compress/gzip"
	json "encoding/json"
	"io"

	context "context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	host "github.com/libp2p/go-libp2p/core/host"
)

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
