package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	examples "github.com/mycontroller-org/esphome_api/examples"
	"github.com/mycontroller-org/esphome_api/pkg/api"
	"google.golang.org/protobuf/proto"
)

func main() {
	client, err := examples.GetClient(handleFuncImpl)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = client.Close() }()

	// subscribe state changes
	if err := client.Send(&api.SubscribeStatesRequest{}); err != nil {
		log.Fatalln(err)
	}

	// wait for a second
	<-time.After(1 * time.Second)

	// take a picture
	if err := client.Send(&api.CameraImageRequest{Single: true}); err != nil {
		log.Fatalln(err)
	}

	// wait 10 seconds
	<-time.After(3 * time.Second)

	// if image received, convert it to jpeg
	if received {
		err = os.WriteFile("camera_image.jpeg", buffer.Bytes(), fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var (
	buffer   = new(bytes.Buffer)
	received = false
)

func handleFuncImpl(msg proto.Message) {
	switch msg := msg.(type) {
	case *api.CameraImageResponse:
		if !received {
			buffer.Write(msg.Data)
			if msg.Done {
				received = true
				fmt.Println("Image received")
			}
		}

	default:
	}
}
