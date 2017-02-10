package main

import pusher "github.com/pusher/pusher-http-go"

func main() {
	client := pusher.Client{
		AppId:   "300550",
		Key:     "c24dabd6884e70c4eafb",
		Secret:  "edc884fc9c00031fa8cb",
		Cluster: "eu",
	}

	data := map[string][]string{}
	data["directions"] = []string{"90", "180", "30"}
	data["headings"] = []string{"30", "30", "30"}

	client.Trigger("test_channel", "my_event", data)
}
