package types

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	EncodePrefix = "BASE64/"
)

type Config struct {
	Active  string         `yaml:"active"`
	Devices []DeviceConfig `yaml:"devices"`
}

type DeviceInfo struct {
	Name            string    `yaml:"name"`
	Model           string    `yaml:"model"`
	MacAddress      string    `yaml:"macAddress"`
	EsphomeVersion  string    `yaml:"esphomeVersion"`
	CompilationTime string    `yaml:"compilationTime"`
	UsesPassword    bool      `yaml:"usesPassword"`
	HasDeepSleep    bool      `yaml:"hasDeepSleep"`
	StatusOn        time.Time `yaml:"statusOn"`
}

func (di *DeviceInfo) Clone() DeviceInfo {
	return DeviceInfo{
		Name:            di.Name,
		Model:           di.Model,
		MacAddress:      di.MacAddress,
		EsphomeVersion:  di.EsphomeVersion,
		CompilationTime: di.CompilationTime,
		UsesPassword:    di.UsesPassword,
		HasDeepSleep:    di.HasDeepSleep,
	}
}

type DeviceConfig struct {
	Address       string        `yaml:"address"`
	Password      string        `yaml:"password"` // encode as base64
	EncryptionKey string        `yaml:"encryptionKey"`
	Timeout       time.Duration `yaml:"timeout"`
	Info          DeviceInfo    `yaml:"info"`
}

func (dc *DeviceConfig) Clone() DeviceConfig {
	return DeviceConfig{
		Address:       dc.Address,
		Password:      dc.Password,
		EncryptionKey: dc.EncryptionKey,
		Timeout:       dc.Timeout,
		Info:          dc.Info.Clone(),
	}
}

// GetPassword decodes and returns the password
func (nc *DeviceConfig) GetPassword() string {
	if strings.HasPrefix(nc.Password, EncodePrefix) {
		password := strings.Replace(nc.Password, EncodePrefix, "", 1)
		decodedPassword, err := base64.StdEncoding.DecodeString(password)
		if err != nil {
			log.Fatal("error on decoding the password", err)
		}
		return string(decodedPassword)
	}
	return nc.Password
}

// EncodePassword encodes and update the password
func (nc *DeviceConfig) EncodePassword() {
	if nc.Password != "" && !strings.HasPrefix(nc.Password, EncodePrefix) {
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(nc.Password))
		nc.Password = fmt.Sprintf("%s%s", EncodePrefix, encodedPassword)
	}
}
