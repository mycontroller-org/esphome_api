package api

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"
)

func Marshal_(message proto.Message) ([]byte, error) {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	bytesPack := make([]byte, len(messageBytes)+20) // message size + message length bytes + message type id bytes

	// keep 0th byte as nil/0
	// refer: https://github.com/esphome/esphome/blob/v1.18.0/esphome/components/api/api_connection.cpp#L50
	// start from the position 1
	index := 1

	// include message bytes length
	index += binary.PutUvarint(bytesPack[index:], uint64(len(messageBytes)))

	// include message type
	index += binary.PutUvarint(bytesPack[index:], TypeID(message))

	// copy message bytes
	copy(bytesPack[index:], messageBytes)
	index += len(messageBytes)

	return bytesPack[:index], nil
}

func Unmarshal_(reader *bufio.Reader) (proto.Message, error) {
	firstByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	if firstByte != 0x00 {
		return nil, errors.New("invalid preamble. should starts with 0x00 byte")
	}

	// get message bytes length
	messageSize, err := binary.ReadUvarint(reader)
	if err != nil {
		return nil, err
	}

	// get message type id
	messageTypeID, err := binary.ReadUvarint(reader)
	if err != nil {
		return nil, err
	}

	// get message bytes
	messageBytes := make([]byte, messageSize)
	_, err = io.ReadFull(reader, messageBytes)
	if err != nil {
		return nil, err
	}

	message := NewMessageByTypeID(messageTypeID)
	if message == nil {
		return nil, fmt.Errorf("api: protocol error: unknown message type %#x", messageTypeID)
	}

	err = proto.Unmarshal(messageBytes, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
