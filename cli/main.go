package main

import (
	_ "github.com/mycontroller-org/esphome_api/cli/command/device"
	rootCmd "github.com/mycontroller-org/esphome_api/cli/command/root"

	clientTY "github.com/mycontroller-org/server/v2/pkg/types/client"
)

func main() {
	rootCmd.Execute(clientTY.NewStdStreams())
}
