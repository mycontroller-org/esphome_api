package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/esphome_api/pkg/connection"
	types "github.com/mycontroller-org/esphome_api/pkg/types"
)

// API version sent in HelloRequest.
const (
	ClientAPIVersionMajor uint32 = 1
	ClientAPIVersionMinor uint32 = 14
)

// Client struct.
type Client struct {
	ID                   string
	conn                 net.Conn
	reader               *bufio.Reader
	stopChan             chan struct{} // closed by requestStop
	stopOnce             sync.Once
	waitMapMutex         sync.RWMutex
	waitMap              map[uint64]chan proto.Message
	lastMessageAt        time.Time
	callBackFunc         types.CallBackFunc
	CommunicationTimeout time.Duration
	apiConn              connection.ApiConnection
	// DisconnectReason from device DisconnectRequest (0 if none).
	DisconnectReason api.DisconnectReason
}

// GetClient returns esphome api client
func GetClient(clientID, address, encryptionKey string, timeout time.Duration, callBackFunc types.CallBackFunc) (*Client, error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}

	// add noop func, if handler not defined
	if callBackFunc == nil {
		callBackFunc = func(msg proto.Message) {}
	}

	apiConn, err := connection.GetConnection(conn, timeout, encryptionKey)
	if err != nil {
		return nil, err
	}

	c := &Client{
		ID:                   clientID,
		conn:                 conn,
		reader:               bufio.NewReader(conn),
		waitMap:              make(map[uint64]chan proto.Message),
		stopChan:             make(chan struct{}),
		callBackFunc:         callBackFunc,
		CommunicationTimeout: timeout,
		apiConn:              apiConn,
	}

	// call handshake, used in encrypted connection
	err = apiConn.Handshake()
	if err != nil {
		return nil, err
	}

	go c.messageReader()
	return c, nil
}

// Close the client
func (c *Client) Close() error {
	_, err := c.SendAndWaitForResponse(&api.DisconnectRequest{}, api.DisconnectResponseTypeID)
	c.requestStop()
	return err
}

func (c *Client) requestStop() {
	c.stopOnce.Do(func() {
		close(c.stopChan)
		_ = c.conn.SetDeadline(time.Now()) // unblock Read
	})
}

// Hello func
func (c *Client) Hello() (*types.HelloResponse, error) {
	response, err := c.SendAndWaitForResponse(&api.HelloRequest{
		ClientInfo:      c.ID,
		ApiVersionMajor: ClientAPIVersionMajor,
		ApiVersionMinor: ClientAPIVersionMinor,
	}, api.HelloResponseTypeID)
	if err != nil {
		return nil, err
	}
	helloResponse, ok := response.(*api.HelloResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type:%T", response)
	}
	return &types.HelloResponse{
		ApiVersionMajor: helloResponse.ApiVersionMajor,
		ApiVersionMinor: helloResponse.ApiVersionMinor,
		ServerInfo:      helloResponse.ServerInfo,
		Name:            helloResponse.Name,
	}, nil
}

// Login uses the legacy API password. Removed in ESPHome 2026.1.0; use Hello() instead.
func (c *Client) Login(password string) error {
	if _, err := c.Hello(); err != nil {
		return err
	}

	// AuthenticationRequest/Response are deprecated in api.proto (ESPHome 2026.1.0+)
	// but still required for pre-2026.1 devices that use api password.
	message, err := c.SendAndWaitForResponse(
		&api.AuthenticationRequest{Password: password}, //nolint:staticcheck // SA1019: legacy password auth
		api.AuthenticationResponseTypeID,
	)
	if err != nil {
		if errors.Is(err, types.ErrCommunicationTimeout) {
			// No reply: treat as modern device without password auth.
			return nil
		}
		return err
	}
	authResponse, ok := message.(*api.AuthenticationResponse) //nolint:staticcheck // SA1019: legacy password auth
	if !ok {
		return fmt.Errorf("invalid response type:%T", message)
	}
	if authResponse.InvalidPassword {
		return types.ErrPassword
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
func (c *Client) DeviceInfo() (*types.DeviceInfo, error) {
	message, err := c.SendAndWaitForResponse(&api.DeviceInfoRequest{}, api.DeviceInfoResponseTypeID)
	if err != nil {
		return nil, err
	}

	info := message.(*api.DeviceInfoResponse)
	di := &types.DeviceInfo{
		UsesPassword:               info.UsesPassword, //nolint:staticcheck // SA1019: still reported by older firmware
		Name:                       info.Name,
		MacAddress:                 info.MacAddress,
		EsphomeVersion:             info.EsphomeVersion,
		CompilationTime:            info.CompilationTime,
		Model:                      info.Model,
		HasDeepSleep:               info.HasDeepSleep,
		ProjectName:                info.ProjectName,
		ProjectVersion:             info.ProjectVersion,
		WebserverPort:              info.WebserverPort,
		Manufacturer:               info.Manufacturer,
		FriendlyName:               info.FriendlyName,
		SuggestedArea:              info.SuggestedArea,
		BluetoothMacAddress:        info.BluetoothMacAddress,
		BluetoothProxyFeatureFlags: info.BluetoothProxyFeatureFlags,
		VoiceAssistantFeatureFlags: info.VoiceAssistantFeatureFlags,
		ApiEncryptionSupported:     info.ApiEncryptionSupported,
		ApiEncryptionProvisionable: info.ApiEncryptionProvisionable,
		ZwaveProxyFeatureFlags:     info.ZwaveProxyFeatureFlags,
		ZwaveHomeId:                info.ZwaveHomeId,
	}

	if area := info.GetArea(); area != nil {
		di.Area = &types.AreaInfo{
			AreaID: area.AreaId,
			Name:   area.Name,
		}
	}
	for _, a := range info.GetAreas() {
		di.Areas = append(di.Areas, types.AreaInfo{AreaID: a.AreaId, Name: a.Name})
	}
	for _, d := range info.GetDevices() {
		di.Devices = append(di.Devices, types.SubDeviceInfo{
			DeviceID: d.DeviceId,
			Name:     d.Name,
			AreaID:   d.AreaId,
		})
	}
	for _, sp := range info.GetSerialProxies() {
		di.SerialProxies = append(di.SerialProxies, types.SerialProxyInfo{
			Name:     sp.Name,
			PortType: int32(sp.PortType),
		})
	}

	return di, nil
}

// SubscribeLogs func
func (c *Client) SubscribeLogs(level types.LogLevel) error {
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

func (c *Client) NoiseEncryptionSetKey(key []byte) (*api.NoiseEncryptionSetKeyResponse, error) {
	message, err := c.SendAndWaitForResponse(&api.NoiseEncryptionSetKeyRequest{
		Key: key,
	}, api.NoiseEncryptionSetKeyResponseTypeID)
	if err != nil {
		return nil, err
	}
	resp, ok := message.(*api.NoiseEncryptionSetKeyResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type:%T", message)
	}
	return resp, nil
}

// messageReader reads message from the node
func (c *Client) messageReader() {
	defer func() { _ = c.conn.Close() }()
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
	message, err := c.apiConn.Read(c.reader)
	if err == nil {
		// ignore empty messages (can happen during encryption handshake)
		if message == nil {
			return nil
		}
		c.lastMessageAt = time.Now()

		c.waitMapMutex.Lock()
		in, found := c.waitMap[api.TypeID(message)]
		c.waitMapMutex.Unlock()
		if found {
			in <- message
		}

		if c.handleInternal(message) {
			return nil
		}
		if c.isExternal(message) && c.callBackFunc != nil {
			c.callBackFunc(message)
		}
	}

	return err
}

// isExternal is true when the message should be passed to callBackFunc.
func (c *Client) isExternal(message proto.Message) bool {
	switch message.(type) {
	case
		*api.HelloResponse,
		*api.AuthenticationResponse, //nolint:staticcheck // SA1019: filter legacy auth replies
		*api.DisconnectResponse,
		*api.PingResponse,
		*api.DeviceInfoResponse,
		*api.NoiseEncryptionSetKeyResponse:
		return false
	}
	return true
}

// handleInternal replies to device protocol requests (ping, disconnect, get time).
func (c *Client) handleInternal(message proto.Message) bool {
	switch msg := message.(type) {
	case *api.DisconnectRequest:
		c.DisconnectReason = msg.Reason
		_ = c.Send(&api.DisconnectResponse{})
		c.requestStop()
		return true

	case *api.PingRequest:
		_ = c.Send(&api.PingResponse{})
		return true

	case *api.GetTimeRequest:
		_ = c.Send(&api.GetTimeResponse{
			EpochSeconds: uint32(time.Now().Unix()),
			Timezone:     localTimezoneName(),
		})
		return true
	}

	return false
}

func localTimezoneName() string {
	if tz := os.Getenv("TZ"); tz != "" {
		return tz
	}
	if time.Local != nil {
		if name := time.Local.String(); name != "" && name != "Local" {
			return name
		}
	}
	_, offset := time.Now().Zone()
	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset
	}
	return fmt.Sprintf("UTC%s%02d:%02d", sign, offset/3600, (offset%3600)/60)
}

func (c *Client) Send(message proto.Message) error {
	return c.apiConn.Write(message)
}

func (c *Client) SendAndWaitForResponse(message proto.Message, messageType uint64) (proto.Message, error) {
	// Register waiter before Send to avoid missing a fast response.
	in := make(chan proto.Message, 1)
	c.waitFor(messageType, in)
	defer c.waitDone(messageType)

	if err := c.Send(message); err != nil {
		return nil, err
	}

	select {
	case msg := <-in:
		return msg, nil
	case <-time.After(c.CommunicationTimeout):
		return nil, types.ErrCommunicationTimeout
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
