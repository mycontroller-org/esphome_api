package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	example "github.com/mycontroller-org/esphome_api/example"
	"github.com/mycontroller-org/esphome_api/pkg/model"
	"google.golang.org/protobuf/proto"
)

func main() {
	var logLevel = flag.Int("log_level", int(model.LogLevelVeryVerbose), "log level (0-6)")
	var monitoringDuration = flag.Duration("monitoring_duration", 30*time.Second, "monitoring duration")
	flag.Parse()

	client, err := example.GetClient(handlerFuncImpl)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	fmt.Println("Listening logs, will terminate in", *monitoringDuration)

	err = client.SubscribeLogs(model.LogLevel(*logLevel))
	if err != nil {
		log.Fatalln(err)
	}

	<-time.After(*monitoringDuration)
}

func handlerFuncImpl(msg proto.Message) {
	log, err := model.GetLogEntry(msg)
	if err != nil {
		fmt.Printf("MSG: %+v\n", msg)
	}
	fmt.Println("Log:", log.String())
}
