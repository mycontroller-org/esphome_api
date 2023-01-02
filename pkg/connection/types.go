package connection

import (
	"bufio"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

type ApiConnection interface {
	Write(message proto.Message) error
	Read(reader *bufio.Reader) (proto.Message, error)
	Handshake() error
}

func GetConnection(conn net.Conn, communicationTimeout time.Duration, encryptionKey string) (ApiConnection, error) {
	if encryptionKey != "" {
		return NewEncryptedConnection(conn, communicationTimeout, encryptionKey)
	}
	return NewPlaintextConnection(conn, communicationTimeout)
}
