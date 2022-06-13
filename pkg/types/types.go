package model

import (
	"fmt"

	"github.com/mycontroller-org/esphome_api/pkg/api"
	"google.golang.org/protobuf/proto"
)

// DeviceInfo struct
type DeviceInfo struct {
	Name            string
	Model           string
	MacAddress      string
	EsphomeVersion  string
	CompilationTime string
	UsesPassword    bool
	HasDeepSleep    bool
}

func (di *DeviceInfo) String() string {
	return fmt.Sprintf("{name: %v, model:%v, mac_address:%v, esphome_version:%v, compilation_time:%v, uses_password:%v, has_deep_sleep:%v}",
		di.Name, di.Model, di.MacAddress, di.EsphomeVersion, di.CompilationTime, di.UsesPassword, di.HasDeepSleep)
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
	SendFailed bool
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
		Level:      LogLevel(entry.Level),
		Message:    entry.Message,
		SendFailed: entry.SendFailed,
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
