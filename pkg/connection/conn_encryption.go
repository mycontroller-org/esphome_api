package connection

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/flynn/noise"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"google.golang.org/protobuf/proto"
)

var (
	ErrHandshakeFailed           = errors.New("esphome_api: handshake failed")
	ErrHandshakeNotDone          = errors.New("esphome_api: handshake not done")
	ErrInvalidEncryptionPreamble = errors.New("esphome_api: invalid preamble. encrypted message should starts with 0x01")
)

// encryption is based on noise protocol
// reference: https://noiseprotocol.org
type EncryptedConnection struct {
	conn                 net.Conn
	CommunicationTimeout time.Duration
	encryptionKey        []byte
	handshakeState       *noise.HandshakeState
	isHandshakeSucceed   bool
	handshakeResponse    []byte
	encryption           *noise.CipherState
	decryption           *noise.CipherState
	writeMutex           sync.Mutex
	readMutex            sync.Mutex
}

func NewEncryptedConnection(conn net.Conn, communicationTimeout time.Duration, encryptionKey string) (ApiConnection, error) {
	// decode base64 key
	_psk, err := base64.StdEncoding.DecodeString(encryptionKey)
	if err != nil {
		return nil, err
	}

	if len(_psk) != 32 {
		return nil, fmt.Errorf("esphome_api: encryption key size should be 32 bytes. received:%d", len(_psk))
	}

	ec := &EncryptedConnection{
		conn:                 conn,
		CommunicationTimeout: communicationTimeout,
		encryptionKey:        _psk,
		isHandshakeSucceed:   false,
	}
	return ec, nil
}

// handshake implemented based on python implementation
// reference: https://github.com/esphome/aioesphomeapi/blob/main/aioesphomeapi/_frame_helper.py#L250
func (ec *EncryptedConnection) Handshake() error {
	// send empty message
	err := ec.writeToTransport(nil)
	if err != nil {
		return err
	}
	// wait for a response
	err = ec.waitForResponse()
	if err != nil {
		return err
	}

	// create handshake state
	_hs, err := noise.NewHandshakeState(
		noise.Config{
			Pattern:               noise.HandshakeNN,
			CipherSuite:           noise.NewCipherSuite(noise.DH25519, noise.CipherChaChaPoly, noise.HashSHA256),
			Initiator:             true,
			Prologue:              []byte("NoiseAPIInit" + "\x00\x00"),
			PresharedKey:          ec.encryptionKey,
			PresharedKeyPlacement: 0,
		},
	)
	if err != nil {
		return err
	}
	ec.handshakeState = _hs

	// create handshake message
	handshakeMessage, _, _, err := ec.handshakeState.WriteMessage(nil, nil)
	if err != nil {
		return err
	}

	// append 0x00 in front of handshake message
	handshakeMessage = append([]byte{0x00}, handshakeMessage...)

	// send handshake message to the transport
	err = ec.writeToTransport(handshakeMessage)
	if err != nil {
		return err
	}

	// wait for handshake response
	// handshake response will be added in ec.handshakeResponse
	err = ec.waitForResponse()
	if err != nil {
		return err
	}

	// if handshake response available, process it and update encrypt and decrypt cipherState
	if len(ec.handshakeResponse) > 0 {
		// discard first byte of handshake response
		_, _encryptCipherState, _decryptCipherState, err := ec.handshakeState.ReadMessage(nil, ec.handshakeResponse[1:])
		if err == nil {
			ec.encryption = _encryptCipherState
			ec.decryption = _decryptCipherState
			ec.isHandshakeSucceed = true
		} else {
			return err
		}
	}

	// return error if handshake not succeed
	if !ec.isHandshakeSucceed {
		return ErrHandshakeFailed
	}

	return nil
}

func (ec *EncryptedConnection) waitForResponse() error {
	_reader := bufio.NewReader(ec.conn)
	_, err := ec.Read(_reader)
	// ignore handshake not done error
	if err != nil && !errors.Is(err, ErrHandshakeNotDone) {
		return err
	}
	return nil
}

func (ec *EncryptedConnection) writeToTransport(data []byte) error {
	if data == nil {
		data = []byte{}
	}

	// preamble + data length (2 bytes fixed)
	header := make([]byte, 3)

	// preamble
	header[0] = 0x01

	// message length
	header[1] = byte(len(data) >> 8 & 0xff)
	header[2] = byte(len(data) & 0xff)

	// append header + data
	packed := append(header, data...)

	// set write lock
	ec.writeMutex.Lock()
	defer ec.writeMutex.Unlock()

	err := ec.conn.SetWriteDeadline(time.Now().Add(ec.CommunicationTimeout))
	if err != nil {
		return err
	}
	_, err = ec.conn.Write(packed)
	return err
}

func (ec *EncryptedConnection) Write(message proto.Message) error {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	// type id (2 bytes) + message length (2 bytes)
	header := make([]byte, 4)

	typeID := api.TypeID(message)
	header[0] = byte(typeID >> 8 & 0xff)
	header[1] = byte(typeID & 0xff)

	header[2] = byte(len(messageBytes) >> 8 & 0xff)
	header[3] = byte(len(messageBytes) & 0xff)

	packed := append(header, messageBytes...)

	// encrypt the data
	encryptedData, err := ec.encrypt(packed)
	if err != nil {
		return err
	}
	return ec.writeToTransport(encryptedData)
}

func (ec *EncryptedConnection) Read(reader *bufio.Reader) (proto.Message, error) {
	// set read lock
	ec.readMutex.Lock()
	defer ec.readMutex.Unlock()

	header := make([]byte, 3) // limit to 3 bytes of header
	_, err := io.ReadFull(reader, header)
	if err != nil {
		return nil, err
	}

	preamble := header[0]
	if preamble != 0x01 {
		return nil, ErrInvalidEncryptionPreamble
	}

	encryptedMessageSize := binary.BigEndian.Uint16(header[1:3])

	// get encrypted message bytes
	encryptedMessageBytes := make([]byte, encryptedMessageSize)
	_, err = io.ReadFull(reader, encryptedMessageBytes)
	if err != nil {
		return nil, err
	}

	// if handshake is not done, these bytes may be a handshake response
	if !ec.isHandshakeSucceed {
		ec.handshakeResponse = encryptedMessageBytes
	}

	// if there is no bytes to decrypt, return from here
	if len(encryptedMessageBytes) == 0 {
		return nil, nil
	}

	// decrypt the message
	decryptedMessageBytes, err := ec.decrypt(encryptedMessageBytes)
	if err != nil {
		return nil, err
	}

	// get type id and message size
	messageTypeID := uint64(binary.BigEndian.Uint16(decryptedMessageBytes[0:2]))
	messageSize := int(binary.BigEndian.Uint16(decryptedMessageBytes[2:4]))

	messageBytes := decryptedMessageBytes[4:]

	if len(messageBytes) != messageSize {
		return nil, fmt.Errorf("esphome_api: message length mismatched. expected:%d, actual:%d", messageSize, len(messageBytes))
	}

	message := api.NewMessageByTypeID(messageTypeID)
	if message == nil {
		return nil, fmt.Errorf("esphome_api: protocol error: unknown message type %#x", messageTypeID)
	}

	err = proto.Unmarshal(messageBytes, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (ec *EncryptedConnection) encrypt(data []byte) ([]byte, error) {
	if ec.isHandshakeSucceed {
		return ec.encryption.Encrypt(nil, nil, data)
	}
	return nil, ErrHandshakeNotDone
}

func (ec *EncryptedConnection) decrypt(data []byte) ([]byte, error) {
	if ec.isHandshakeSucceed {
		return ec.decryption.Decrypt(nil, nil, data)
	}
	return nil, ErrHandshakeNotDone
}
