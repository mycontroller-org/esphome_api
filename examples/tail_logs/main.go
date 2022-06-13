package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	examples "github.com/mycontroller-org/esphome_api/examples"
	types "github.com/mycontroller-org/esphome_api/pkg/types"
	"google.golang.org/protobuf/proto"
)

func main() {
	var logLevel = flag.Int("log_level", int(types.LogLevelVeryVerbose), "log level (0-6)")
	var monitoringDuration = flag.Duration("monitoring_duration", 30*time.Second, "monitoring duration")
	flag.Parse()

	client, err := examples.GetClient(handlerFuncImpl)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	fmt.Println("Listening logs, will terminate in", *monitoringDuration)

	err = client.SubscribeLogs(types.LogLevel(*logLevel))
	if err != nil {
		log.Fatalln(err)
	}

	<-time.After(*monitoringDuration)
}

func handlerFuncImpl(msg proto.Message) {
	log, err := types.GetLogEntry(msg)
	if err != nil {
		fmt.Printf("MSG: %+v\n", msg)
	}
	fmt.Println("Log:", log.String())
}
