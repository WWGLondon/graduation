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
		fmt.Printf("Got Message: %#v\n", message)
	}
}

// # Team 3
// 1. Build a pusher URL
// 2. Consruct a client to connect to pusher
// 3. Listen to pusher messages and build a data structure
// 4. Pass data to team 4
