package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/model"
)

// Error types
var (
	ErrPassword             = errors.New("esphome_api: invalid password")
	ErrCommunicationTimeout = errors.New("esphome_api: communication timeout")
)

// Client struct.
type Client struct {
	ID                   string
	conn                 net.Conn
	reader               *bufio.Reader
	stopChan             chan bool
	waitMapMutex         sync.RWMutex
	waitMap              map[uint64]chan proto.Message
	lastMessageAt        time.Time
	handlerFunc          func(proto.Message)
	CommunicationTimeout time.Duration
}

// Init func
func Init(clientID, addr string, timeout time.Duration, handlerFunc func(proto.Message)) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	// add noop func, if handler not defined
	if handlerFunc == nil {
		handlerFunc = func(msg proto.Message) {}
	}

	c := &Client{
		ID:                   clientID,
		conn:                 conn,
		reader:               bufio.NewReader(conn),
		waitMap:              make(map[uint64]chan proto.Message),
		stopChan:             make(chan bool),
		handlerFunc:          handlerFunc,
		CommunicationTimeout: timeout,
	}
	go c.messageReader()
	return c, nil
}

// Close the client
func (c *Client) Close() error {
	_, err := c.SendAndWaitForResponse(&api.DisconnectRequest{}, api.DisconnectResponseTypeID)
	select {
	case c.stopChan <- true:
	default:
	}
	return err
}

// Hello func
func (c *Client) Hello() (*model.HelloResponse, error) {
	response, err := c.SendAndWaitForResponse(&api.HelloRequest{
		ClientInfo: c.ID,
	}, api.HelloResponseTypeID)
	if err != nil {
		return nil, err
	}
	helloResponse, ok := response.(*api.HelloResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type:%T", response)
	}
	return &model.HelloResponse{
		ApiVersionMajor: helloResponse.ApiVersionMajor,
		ApiVersionMinor: helloResponse.ApiVersionMinor,
		ServerInfo:      helloResponse.ServerInfo,
	}, nil
}

// Login func
func (c *Client) Login(password string) error {
	_, err := c.Hello()
	if err != nil {
		return err
	}

	message, err := c.SendAndWaitForResponse(&api.ConnectRequest{
		Password: password,
	}, api.ConnectResponseTypeID)
	if err != nil {
		return err
	}
	connectResponse := message.(*api.ConnectResponse)
	if connectResponse.InvalidPassword {
		return ErrPassword
	}

	return nil
}

// Ping func
func (c *Client) Ping() error {
	_, err := c.SendAndWaitForResponse(&api.PingRequest{}, api.PingResponseTypeID)
	return err
}

// SubscribeStates func
func (c *Client) SubscribeStates() error {
	if err := c.Send(&api.SubscribeStatesRequest{}); err != nil {
		return err
	}
	return nil
}

// LastMessage returns the time of the last message received.
func (c *Client) LastMessageAt() time.Time {
	return c.lastMessageAt
}

// DeviceInfo queries the ESPHome device information.
func (c *Client) DeviceInfo() (*model.DeviceInfo, error) {
	message, err := c.SendAndWaitForResponse(&api.DeviceInfoRequest{}, api.DeviceInfoResponseTypeID)
	if err != nil {
		return nil, err
	}

	info := message.(*api.DeviceInfoResponse)
	return &model.DeviceInfo{
		UsesPassword:    info.UsesPassword,
		Name:            info.Name,
		MacAddress:      info.MacAddress,
		EsphomeVersion:  info.EsphomeVersion,
		CompilationTime: info.CompilationTime,
		Model:           info.Model,
		HasDeepSleep:    info.HasDeepSleep,
	}, nil
}

// SubscribeLogs func
func (c *Client) SubscribeLogs(level model.LogLevel) error {
	if err := c.Send(&api.SubscribeLogsRequest{
		Level: api.LogLevel(level),
	}); err != nil {
		return err
	}

	return nil
}

// ListEntities func
func (c *Client) ListEntities() error {
	return c.Send(&api.ListEntitiesRequest{})
}

// messageReader reads message from the node
func (c *Client) messageReader() {
	defer c.conn.Close()
	for {
		select {
		case <-c.stopChan:
			return

		default:
			if err := c.getMessage(); err != nil {
				return
			}
		}
	}
}

func (c *Client) getMessage() error {
	var message proto.Message
	message, err := api.ReadMessage(c.reader)
	if err == nil {
		c.lastMessageAt = time.Now()
		// check waiting map
		c.waitMapMutex.Lock()
		in, found := c.waitMap[api.TypeID(message)]
		c.waitMapMutex.Unlock()
		if found {
			in <- message
		}

		// forward to other parties
		if c.handleInternal(message) {
			return nil
		} else if c.isExternal(message) {
			if c.handlerFunc != nil {
				c.handlerFunc(message)
				return nil
			}
		}
	}

	return err
}

func (c *Client) isExternal(message proto.Message) bool {
	switch message.(type) {
	case
		*api.PingResponse,
		*api.HelloResponse,
		*api.ConnectResponse,
		*api.DeviceInfoResponse,
		*api.DisconnectResponse:
		return false
	}
	return true
}

func (c *Client) handleInternal(message proto.Message) bool {
	switch message.(type) {
	case *api.DisconnectRequest:
		_ = c.Send(&api.DisconnectResponse{})
		c.Close()
		return true

	case *api.PingRequest:
		_ = c.Send(&api.PingResponse{})
		return true

	case *api.HelloRequest:
		_ = c.Send(&api.HelloResponse{})
		return true

	case *api.ConnectRequest:
		_ = c.Send(&api.ConnectResponse{})
		return true

	}

	return false
}

func (c *Client) Send(message proto.Message) error {
	packed, err := api.Marshal(message)
	if err != nil {
		return err
	}
	if err = c.conn.SetWriteDeadline(time.Now().Add(c.CommunicationTimeout)); err != nil {
		return err
	}
	if _, err = c.conn.Write(packed); err != nil {
		return err
	}
	if err = c.conn.SetWriteDeadline(time.Time{}); err != nil {
		return err
	}
	return nil
}

func (c *Client) SendAndWaitForResponse(message proto.Message, messageType uint64) (proto.Message, error) {
	if err := c.Send(message); err != nil {
		return nil, err
	}
	return c.waitForMessage(messageType)
}

func (c *Client) waitForMessage(messageType uint64) (proto.Message, error) {
	in := make(chan proto.Message, 1)
	c.waitFor(messageType, in)
	defer c.waitDone(messageType)

	select {
	case message := <-in:
		return message, nil
	case <-time.After(c.CommunicationTimeout):
		return nil, ErrCommunicationTimeout
	}
}

func (c *Client) waitFor(messageType uint64, in chan proto.Message) {
	c.waitMapMutex.Lock()
	defer c.waitMapMutex.Unlock()

	other, waiting := c.waitMap[messageType]
	if waiting {
		other <- nil
		close(other)
	}
	c.waitMap[messageType] = in
}

func (c *Client) waitDone(messageType uint64) {
	c.waitMapMutex.Lock()
	defer c.waitMapMutex.Unlock()
	delete(c.waitMap, messageType)
}
