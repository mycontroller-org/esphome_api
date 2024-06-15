package root

import (
	"fmt"

	"github.com/mycontroller-org/esphome_api/cli/version"
	"github.com/spf13/cobra"
)

func init() {
	AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:  "version",
	Long: "Prints version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get().String())
	},
}
