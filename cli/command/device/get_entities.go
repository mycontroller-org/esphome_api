package device

import (
	"fmt"
	"strings"
	"time"

	rootCmd "github.com/mycontroller-org/esphome_api/cli/command/root"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"github.com/mycontroller-org/server/v2/pkg/utils/convertor"
	filterUtils "github.com/mycontroller-org/server/v2/pkg/utils/filter_sort"
	"github.com/mycontroller-org/server/v2/pkg/utils/printer"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var (
	entitiesTimeout time.Duration
)

func init() {
	getCmd.AddCommand(getEntitiesCmd)
	getEntitiesCmd.Flags().DurationVar(&entitiesTimeout, "timeout", 5*time.Second, "Timeout to wait for entities")
}

var getEntitiesCmd = &cobra.Command{
	Use:     "entity",
	Short:   "Lists available entities",
	Aliases: []string{"entities"},
	Example: `  # lists available entities
  esphomectl get entities

  # list entities with a timeout
  esphomectl get entities --timeout 10s
`,
	Run: func(cmd *cobra.Command, args []string) {

		entitiesCollectionDone := false
		entities := map[string][]interface{}{}
		collectEntities := func(msg proto.Message) {
			switch entity := msg.(type) {
			case *api.ListEntitiesDoneResponse:
				entitiesCollectionDone = true

			default:
				_, _deviceClass, err := filterUtils.GetValueByKeyPath(entity, "deviceClass")
				if err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), "error:", err)
					return
				}
				deviceClass := convertor.ToString(_deviceClass)
				_sensors, found := entities[deviceClass]
				if !found {
					_sensors = make([]interface{}, 0)
				}
				_sensors = append(_sensors, entity)
				entities[deviceClass] = _sensors
			}
		}

		client, err := rootCmd.GetActiveClient(collectEntities)
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "error:", err.Error())
			return
		}

		err = client.ListEntities()
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "error:", err.Error())
			return
		}

		ticker := time.NewTicker(200 * time.Millisecond)
		timeoutTime := time.Now().Add(entitiesTimeout)
		for range ticker.C {
			if entitiesCollectionDone || time.Now().Before(timeoutTime) {
				break
			}
		}

		if len(entities) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No resource found")
			return
		}

		for k, _sensors := range entities {
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintln(cmd.OutOrStdout(), strings.ToUpper(k))

			switch k {
			case "light":
				headers := []printer.Header{
					{Title: "name", ValuePath: "name"},
					{Title: "object id", ValuePath: "objectId"},
					{Title: "key", ValuePath: "key"},
					{Title: "unique id", ValuePath: "uniqueId"},
					{Title: "effects", ValuePath: "effects"},
					{Title: "icon", ValuePath: "icon"},
				}
				printer.Print(cmd.OutOrStdout(), headers, _sensors, rootCmd.HideHeader, rootCmd.OutputFormat, rootCmd.Pretty)

			default:
				headers := []printer.Header{
					{Title: "name", ValuePath: "name"},
					{Title: "object id", ValuePath: "objectId"},
					{Title: "key", ValuePath: "key"},
					{Title: "unique id", ValuePath: "uniqueId"},
					{Title: "device class", ValuePath: "deviceClass"},
				}
				printer.Print(cmd.OutOrStdout(), headers, _sensors, rootCmd.HideHeader, rootCmd.OutputFormat, rootCmd.Pretty)

			}

		}

	},
}
