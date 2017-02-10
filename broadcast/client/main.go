package main

import (
	"fmt"
	"log"

	"github.com/WWGLondon/graduation/broadcast/client/pusher"
)

var appID = "c24dabd6884e70c4eafb"

func main() {
	origin := "http://localhost/"

	options := pusher.Options{
		Cluster:       "eu",
		Client:        "go",
		ClientVersion: "0.1",
		APIVersion:    "7",
		Protocol:      "wss",
		AppId:         appID,
	}
	wsURL := pusher.NewURL(options)
	connection := pusher.Connection{}

	fmt.Println(wsURL.String())

	err := connection.Connect(wsURL, origin)
	if err != nil {
		log.Fatal(err)
	}

	c, err := connection.Subscribe("test_channel")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening to channel")
	for message := range c {
		fmt.Println("Got Message ", message)
	}
}
