package erlgo

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

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
