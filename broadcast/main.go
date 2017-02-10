package broadcast

import pusher "github.com/pusher/pusher-http-go"

func main() {
	client := pusher.Client{
		AppId:   "300550",
		Key:     "c24dabd6884e70c4eafb",
		Secret:  "edc884fc9c00031fa8cb",
		Cluster: "eu",
	}

	data := map[string][]string{}
	data["speed"] = []string{"walking", "running", "running", "crawling"}
	data["direction"] = []string{"North", "East", "South", "West"}

	client.Trigger("test_channel", "my_event", data)
}
