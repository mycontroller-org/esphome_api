package device

import (
	"github.com/spf13/cobra"
)

func init() {
	setCmd.AddCommand(setEntitiesCmd)
}

var setEntitiesCmd = &cobra.Command{
	Use:   "entity",
	Short: "Updates entity value",
	Example: `  # lists available entities
  esphomectl set entities unique_id --payload test=on
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// _client, err := rootCmd.GetActiveClient(nil)
		// if err != nil {
		// 	fmt.Fprintln(cmd.ErrOrStderr(), "error:", err.Error())
		// 	return
		// }
		//
		// var request proto.Message

	},
}
