package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
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
	defer client.Close()

	// subscribe state changes
	client.Send(&api.SubscribeStatesRequest{})

	// wait for a second
	<-time.After(1 * time.Second)

	// take a picture
	client.Send(&api.CameraImageRequest{Single: true})

	// wait 10 seconds
	<-time.After(3 * time.Second)

	// if image received, convert it to jpeg
	if received {
		err = ioutil.WriteFile("camera_image.jpeg", buffer.Bytes(), fs.ModePerm)
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
