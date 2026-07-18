package model

import (
	"errors"
	"fmt"

	"github.com/mycontroller-org/esphome_api/pkg/api"
	"google.golang.org/protobuf/proto"
)

// Error types
var (
	ErrPassword              = errors.New("esphome_api: invalid password")
	ErrCommunicationTimeout  = errors.New("esphome_api: communication timeout")
	ErrConnRequireEncryption = errors.New("esphome_api: connection requires encryption")
)

// call back function used to report received messages
type CallBackFunc func(proto.Message)

type AreaInfo struct {
	AreaID uint32
	Name   string
}

type SubDeviceInfo struct {
	DeviceID uint32
	Name     string
	AreaID   uint32
}

type SerialProxyInfo struct {
	Name     string
	PortType int32
}

// DeviceInfo struct
type DeviceInfo struct {
	Name                       string
	Model                      string
	MacAddress                 string
	EsphomeVersion             string
	CompilationTime            string
	UsesPassword               bool // removed in ESPHome 2026.1.0
	HasDeepSleep               bool
	ProjectName                string
	ProjectVersion             string
	WebserverPort              uint32
	Manufacturer               string
	FriendlyName               string
	SuggestedArea              string
	BluetoothMacAddress        string
	BluetoothProxyFeatureFlags uint32
	VoiceAssistantFeatureFlags uint32
	ApiEncryptionSupported     bool
	ApiEncryptionProvisionable bool
	ZwaveProxyFeatureFlags     uint32
	ZwaveHomeId                uint32
	Area                       *AreaInfo
	Areas                      []AreaInfo
	Devices                    []SubDeviceInfo
	SerialProxies              []SerialProxyInfo
}

func (di *DeviceInfo) String() string {
	return fmt.Sprintf("{name: %v, friendly_name: %v, model:%v, manufacturer:%v, mac_address:%v, esphome_version:%v, compilation_time:%v, uses_password:%v, has_deep_sleep:%v, project:%v/%v, webserver_port:%v, suggested_area:%v, bt_mac:%v, api_encryption_supported:%v}",
		di.Name, di.FriendlyName, di.Model, di.Manufacturer, di.MacAddress, di.EsphomeVersion, di.CompilationTime, di.UsesPassword, di.HasDeepSleep, di.ProjectName, di.ProjectVersion, di.WebserverPort, di.SuggestedArea, di.BluetoothMacAddress, di.ApiEncryptionSupported)
}

// LogLevel type
type LogLevel int32

// log levels
const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug // default
	LogLevelVerbose
	LogLevelVeryVerbose
)

// LogEntry of a message
type LogEntry struct {
	Level      LogLevel
	Tag        string
	Message    string
	SendFailed bool // always false; field removed from API
}

func (le *LogEntry) String() string {
	return fmt.Sprintf("{level: %v, tag:%v, send_failed:%v, message:[%v]}",
		le.Level, le.Tag, le.SendFailed, le.Message)
}

func GetLogEntry(msg proto.Message) (*LogEntry, error) {
	entry, ok := msg.(*api.SubscribeLogsResponse)
	if !ok {
		return nil, fmt.Errorf("received invalid data type:%T", msg)
	}
	log := LogEntry{
		Level:   LogLevel(entry.Level),
		Message: string(entry.Message),
	}
	return &log, nil
}

type HelloResponse struct {
	ApiVersionMajor uint32
	ApiVersionMinor uint32
	ServerInfo      string
	Name            string
}

func (hr *HelloResponse) String() string {
	return fmt.Sprintf("{name: %v, api_version_major: %v, api_version_minor:%v, server_info:%v}",
		hr.Name, hr.ApiVersionMajor, hr.ApiVersionMinor, hr.ServerInfo)
}
