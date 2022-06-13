package main

import (
	"fmt"
	"log"

	examples "github.com/mycontroller-org/esphome_api/examples"
)

func main() {
	client, err := examples.GetClient(nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	err = client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	hr, err := client.Hello()
	if err != nil {
		log.Fatalln(err)
	}

	di, err := client.DeviceInfo()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("HelloResponse:", hr.String())
	fmt.Println("DeviceInfo:", di.String())
}
