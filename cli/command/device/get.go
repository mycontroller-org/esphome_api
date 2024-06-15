package device

import (
	rootCmd "github.com/mycontroller-org/esphome_api/cli/command/root"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List resources",
}
