package root

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	cliTY "github.com/mycontroller-org/esphome_api/cli/types"
	"github.com/mycontroller-org/esphome_api/pkg/client"
	TY "github.com/mycontroller-org/esphome_api/pkg/types"
	clientTY "github.com/mycontroller-org/server/v2/pkg/types/client"
	printer "github.com/mycontroller-org/server/v2/pkg/utils/printer"
	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ENV_PREFIX       = "ESPHOME"
	CONFIG_FILE_NAME = ".esphomectl"
	CONFIG_FILE_EXT  = "yaml"
	CONFIG_FILE_ENV  = "ESPHOMECTL_CONFIG"
)

var (
	cfgFile   string
	CONFIG    *cliTY.Config      // keeps device details
	ioStreams clientTY.IOStreams // read and write to this stream

	HideHeader   bool
	Pretty       bool
	OutputFormat string

	rootCliLong = `esphome cli Client
  
This client helps you to control your esphome devices from the command line.
`
)

func AddCommand(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}

var rootCmd = &cobra.Command{
	Use:   "esphomectl",
	Short: "esphomectl",
	Long:  rootCliLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.SetOut(ioStreams.Out)
		cmd.SetErr(ioStreams.ErrOut)
	},
}

func init() {
	CONFIG = &cliTY.Config{}

	cobra.OnInitialize(loadConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.esphomectl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", printer.OutputConsole, "output format. options: yaml, json, console, wide")
	rootCmd.PersistentFlags().BoolVar(&HideHeader, "hide-header", false, "hides the header on the console output")
	rootCmd.PersistentFlags().BoolVar(&Pretty, "pretty", false, "JSON pretty print")
}

func GetActiveClient(callBackFunc TY.CallBackFunc) (*client.Client, error) {
	if CONFIG.Active == "" {
		return nil, errors.New("no device configured")
	}
	cfg := GetDevice(CONFIG.Active)
	if cfg == nil {
		return nil, fmt.Errorf("device[%s] configuration is not available", CONFIG.Active)
	}
	return GetClient(cfg, callBackFunc)
}

func GetClient(cfg *cliTY.DeviceConfig, callBackFunc TY.CallBackFunc) (*client.Client, error) {
	_client, err := client.GetClient("mc_esphome_cli", cfg.Address, cfg.EncryptionKey, cfg.Timeout, callBackFunc)
	if err != nil {
		return nil, err
	}
	if cfg.GetPassword() != "" {
		err = _client.Login(cfg.GetPassword())
		if err != nil {
			return nil, err
		}
	}

	return _client, nil
}

func Execute(streams clientTY.IOStreams) {
	ioStreams = streams
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(ioStreams.ErrOut, err)
		os.Exit(1)
	}
}

func WriteConfigFile() {
	if cfgFile == "" {
		return
	}
	if CONFIG == nil {
		CONFIG = &cliTY.Config{}
	}

	configBytes, err := yaml.Marshal(CONFIG)
	if err != nil {
		fmt.Fprintf(ioStreams.ErrOut, "error on config file marshal. error:[%s]\n", err.Error())
	}
	err = os.WriteFile(cfgFile, configBytes, os.ModePerm)
	if err != nil {
		fmt.Fprintf(ioStreams.ErrOut, "error on writing config file to disk, filename:%s, error:[%s]\n", cfgFile, err.Error())
	}
}

func loadConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else if os.Getenv(CONFIG_FILE_ENV) != "" {
		cfgFile = os.Getenv(CONFIG_FILE_ENV)
	} else {
		// Find home directory.initConfig
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".myc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(CONFIG_FILE_NAME)
		viper.SetConfigType(CONFIG_FILE_EXT)

		cfgFile = filepath.Join(home, fmt.Sprintf("%s.%s", CONFIG_FILE_NAME, CONFIG_FILE_EXT))
	}

	viper.SetEnvPrefix(ENV_PREFIX)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		err = viper.Unmarshal(&CONFIG)
		if err != nil {
			fmt.Fprint(ioStreams.ErrOut, "error on unmarshal of config\n", err)
		}
	}
}

func GetDevice(address string) *cliTY.DeviceConfig {
	for _, device := range CONFIG.Devices {
		if device.Address == address {
			_device := device.Clone()
			return &_device
		}
	}
	return nil
}

func RemoveDevice(address string) {
	devices := make([]cliTY.DeviceConfig, 0)
	for _, device := range CONFIG.Devices {
		if device.Address == address {
			continue
		}
		devices = append(devices, device)
	}
	CONFIG.Devices = devices
	CONFIG.Active = ""
}

func AddDevice(newDevice *cliTY.DeviceConfig) {
	newDevice.EncodePassword()
	devices := make([]cliTY.DeviceConfig, 0)
	for _, device := range CONFIG.Devices {
		if device.Address == newDevice.Address {
			continue
		}
		devices = append(devices, device)
	}

	// add current device
	devices = append(devices, *newDevice)
	CONFIG.Devices = devices
	CONFIG.Active = newDevice.Address
}
