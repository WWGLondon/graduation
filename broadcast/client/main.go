// Team 3 solution from https://github.com/Johanna-hub/graduation/blob/master/broadcast/client/main.go
package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/WWGLondon/graduation/broadcast/client/pusher"
)

func main() {
	options := pusher.Options{
		Cluster:       "eu",
		Client:        "go",
		ClientVersion: "0.1",
		APIVersion:    "7",
		Protocol:      "wss",
		AppId:         "APP_ID", // update me with actual Pusher APP ID
	}
	pusherURL := pusher.NewURL(options)

	pusherConnect := pusher.Connection{}
	err := pusherConnect.Connect(pusherURL, "http://localhost/")
	if err != nil {
		fmt.Println(err)
	}

	events, err := pusherConnect.Subscribe("my-channel")
	fmt.Println(events)
	if err != nil {
		fmt.Println(err)
	}

	for m := range events {
		// Update the post URL with actual address where you are posting this to
		resp, err := http.Post("http://172.16.14.244:7001/input", "application/json", bytes.NewBufferString(m.Data))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp)
	}
}
