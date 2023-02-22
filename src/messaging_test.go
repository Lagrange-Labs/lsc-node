package main

import (
	"bytes"
	"compress/gzip"
	"reflect"
	"testing"
)

func TestCompressMessage(t *testing.T) {
	t.Logf("TestCompressMessage started")
	message := []byte("Hello World!")
	compressedMessage := compressMessage(message)

	// Check that the compressed message is not nil
	if compressedMessage == nil {
		t.Error("Expected compressed message to not be nil")
	}

	// Check that the compressed message is not the same as the original message
	if reflect.DeepEqual(compressedMessage, message) {
		t.Error("Expected compressed message to be different from the original message")
	}

	// Check that the compressed message can be decompressed to the original message
	decompressedMessage, err := decompressMessage(compressedMessage)
	if err != nil {
		t.Error("Failed to decompress message:", err)
	}
	if !bytes.Equal(decompressedMessage, message) {
		t.Error("Decompressed message does not match the original message")
	}
	t.Logf("TestCompressMessage completed successfully")
}

func TestDecompressMessage(t *testing.T) {
	t.Logf("TestDecompressMessage started")
	message := []byte("Hello World!")
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(message); err != nil {
		t.Error("Failed to write message to gzip writer:", err)
	}
	if err := gz.Flush(); err != nil {
		t.Error("Failed to flush gzip writer:", err)
	}
	if err := gz.Close(); err != nil {
		t.Error("Failed to close gzip writer:", err)
	}
	compressedMessage := b.Bytes()

	// Check that the decompressed message can be obtained from the compressed message
	decompressedMessage, err := decompressMessage(compressedMessage)
	if err != nil {
		t.Error("Failed to decompress message:", err)
	}
	if !bytes.Equal(decompressedMessage, message) {
		t.Error("Decompressed message does not match the original message")
	}

	// Check that decompressing an invalid message returns an error
	_, err = decompressMessage([]byte{0x00, 0x01, 0x02, 0x03})
	if err == nil {
		t.Error("Expected decompressing an invalid message to return an error")
	}
	t.Logf("TestDecompressMessage completed successfully")
}
