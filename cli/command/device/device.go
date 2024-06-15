package device

import (
	"fmt"
	"strings"

	rootCmd "github.com/mycontroller-org/esphome_api/cli/command/root"
	"github.com/mycontroller-org/server/v2/pkg/utils/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deviceContextCmd)
	rootCmd.AddCommand(getDevicesCmd)
}

var deviceContextCmd = &cobra.Command{
	Use:   "device",
	Short: "Switch or set a device",
	Example: `  # set a node
  esphomectl device my-device-1

  # get the active device
  esphomectl device
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if rootCmd.CONFIG.Active == "" {
				fmt.Fprintln(cmd.OutOrStdout(), "No resource found")
				return
			}
			fmt.Fprintf(cmd.ErrOrStderr(), "Active node '%s'\n", rootCmd.CONFIG.Active)
			return
		}
		rootCmd.CONFIG.Active = strings.TrimSpace(args[0])
		client, err := rootCmd.GetActiveClient(nil)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "Error on login", err)
			return
		}
		if client != nil {
			rootCmd.WriteConfigFile()
			fmt.Fprintf(cmd.OutOrStdout(), "Switched to '%s'\n", rootCmd.CONFIG.Active)
		}
	},
}

var getDevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Display configured devices",
	Example: `  # display configured devices
  esphomectl devices
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(rootCmd.CONFIG.Devices) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No resource found")
			return
		}
		headers := []printer.Header{
			{Title: "address", ValuePath: "address"},
			{Title: "name", ValuePath: "info.name"},
			{Title: "model", ValuePath: "info.model"},
			{Title: "mac address", ValuePath: "info.macAddress"},
			{Title: "version", ValuePath: "info.esphomeVersion"},
			{Title: "compilation time", ValuePath: "info.compilationTime"},
			{Title: "uses password", ValuePath: "info.usesPassword"},
			{Title: "has deep sleep", ValuePath: "info.hasDeepSleep"},
			{Title: "timeout", ValuePath: "timeout", IsWide: true},
			{Title: "status on", ValuePath: "info.statusOn", DisplayStyle: printer.DisplayStyleRelativeTime},
		}
		data := make([]interface{}, 0)
		for _, device := range rootCmd.CONFIG.Devices {
			data = append(data, device)
		}
		printer.Print(cmd.OutOrStdout(), headers, data, rootCmd.HideHeader, rootCmd.OutputFormat, rootCmd.Pretty)
	},
}
