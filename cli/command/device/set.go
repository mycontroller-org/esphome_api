package device

import (
	rootCmd "github.com/mycontroller-org/esphome_api/cli/command/root"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "update resource state",
}
