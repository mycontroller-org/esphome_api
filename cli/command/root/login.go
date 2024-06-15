package root

import (
	"fmt"
	"time"

	cliTY "github.com/mycontroller-org/esphome_api/cli/types"
	"github.com/spf13/cobra"
)

var (
	devicePassword      string
	deviceEncryptionKey string
	deviceTimeout       time.Duration
)

func init() {
	AddCommand(loginCmd)
	loginCmd.Flags().StringVar(&devicePassword, "password", "", "Password to login into esphome device")
	loginCmd.Flags().StringVar(&deviceEncryptionKey, "encryption-key", "", "Encryption key to login into esphome device")
	loginCmd.Flags().DurationVar(&deviceTimeout, "timeout", 10*time.Second, "esphome device communication timeout")

	AddCommand(logoutCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a esphome device",
	Example: `  # login to esphome device without password and encryption key
  esphomectl login my_esphome.local:6053

  # login to esphome device with password
  esphomectl login my_esphome.local:6053 --password my_secret

  # login to esphome device with encryption key
  esphomectl login my_esphome.local:6053 --encryption-key my_encryption_key
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deviceCfg := &cliTY.DeviceConfig{
			Address:       args[0],
			Password:      devicePassword,
			EncryptionKey: deviceEncryptionKey,
			Timeout:       deviceTimeout,
		}

		_client, err := GetClient(deviceCfg, nil)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "error on login", err)
			return
		}
		if _client != nil {
			deviceInfo, err := _client.DeviceInfo()
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), "error on getting device information", err)
				return
			}
			// update device info
			deviceCfg.Info = cliTY.DeviceInfo{
				Name:            deviceInfo.Name,
				Model:           deviceInfo.Model,
				MacAddress:      deviceInfo.MacAddress,
				EsphomeVersion:  deviceInfo.EsphomeVersion,
				CompilationTime: deviceInfo.CompilationTime,
				UsesPassword:    deviceInfo.UsesPassword,
				HasDeepSleep:    deviceInfo.HasDeepSleep,
				StatusOn:        time.Now(),
			}
			AddDevice(deviceCfg)
			WriteConfigFile()

			fmt.Fprintln(cmd.OutOrStdout(), "Login successful.")
			fmt.Fprintf(cmd.OutOrStdout(), "%+v\n", deviceInfo)
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from a esphome device",
	Example: `  # logout from a esphome device
  esphomectl logout

  # logout from esphome devices
  esphomectl logout my_device_1:6053 my_device_2:6053`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && CONFIG.Active == "" {
			fmt.Fprintln(cmd.ErrOrStderr(), "There is no active device information.")
			return
		}

		// remove given devices
		for _, address := range args {
			RemoveDevice(address)
		}
		WriteConfigFile()

		fmt.Fprintln(cmd.OutOrStdout(), "Logout successful.")
	},
}
