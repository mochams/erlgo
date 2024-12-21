// Package erlgo provides functionality for communicating with Erlang ports.
// It implements the Erlang port protocol which consists of length-prefixed
// messages where each message is preceded by its length as a 4-byte integer
// in big-endian order.

package erlgo

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

// ReadFromErlang reads a message from Erlang via stdin.
// The message follows the Erlang port protocol format:
// - 4-byte length prefix in big-endian order
// - followed by the actual message content
//
// The function returns the message content as a byte slice and any error
// encountered during reading. Possible errors include:
// - EOF if the port is closed
// - I/O errors from reading stdin
// - Message length exceeding available memory
//
// Example:
//
//	messageBytes, err := erlgo.ReadFromErlang()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Process messageBytes...
func ReadFromErlang() ([]byte, error) {
	// Create a new buffered reader from stdin
	reader := bufio.NewReader(os.Stdin)

	// Read the 4-byte length prefix
	lengthBytes := make([]byte, 4)
	_, err := io.ReadFull(reader, lengthBytes)
	if err != nil {
		return nil, err
	}

	// Convert the length prefix to a uint32 (big-endian)
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read the JSON message based on the length
	messageBytes := make([]byte, length)
	_, err = io.ReadFull(reader, messageBytes)
	if err != nil {
		return nil, err
	}

	return messageBytes, nil
}

// WriteToErlang writes a message to Erlang via stdout following the port protocol.
// The message is automatically prefixed with its length as a 4-byte integer
// in big-endian order.
//
// Parameters:
// - messageBytes: The content to be sent to Erlang
//
// Returns an error if the write operation fails. Possible errors include:
// - I/O errors from writing to stdout
// - System errors if stdout is closed
//
// Example:
//
//	message := []byte("hello")
//	err := erlgo.WriteToErlang(message)
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteToErlang(messageBytes []byte) error {
	// Create length bytes in big-endian order
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(len(messageBytes)))

	// Write length followed by message to stdout
	writer := os.Stdout
	_, err := writer.Write(lengthBytes)
	if err != nil {
		return err
	}
	_, err = writer.Write(messageBytes)
	if err != nil {
		return err
	}

	return nil
}
