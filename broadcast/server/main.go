package main

import (
	"fmt"

	pusher "github.com/pusher/pusher-http-go"
)

func main() {
	client := pusher.Client{
		AppId:   "300550",
		Key:     "c24dabd6884e70c4eafb",
		Secret:  "edc884fc9c00031fa8cb",
		Secure:  true,
		Cluster: "eu",
	}

	data := map[string][]string{}
	data["speed"] = []string{"walking", "running", "running", "crawling"}
	data["direction"] = []string{"North", "East", "South", "West"}

	client.Trigger("test_channel", "my_event", data)

	fmt.Println("Sent Message")
}

// # Team 2
// 1. Identify structure with team 1
// 2. Convert team 1 structure into pusher format
// 3. Create a client
// 4. Send the data to pusher
