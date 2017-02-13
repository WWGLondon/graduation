// Team 2 solution from @vannio
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pusher/pusher-http-go"
)

type Command struct {
	direction string
	mode      string
	order     int
}

type Commands []Command

var receivedCommands Commands

func main() {
	http.HandleFunc("/senddata", sendData)
	http.ListenAndServe(":9000", http.DefaultServeMux)
}

func convertData(anything Command) map[string]string {
	data := map[string]string{}
	data["speed"] = anything.mode
	data["direction"] = anything.direction
	return data
}

func sendData(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &receivedCommands)
	if err != nil {
		log.Fatal(err)
	}
	sendToPusher(receivedCommands)
	fmt.Println("hello")
}

func sendToPusher(commands Commands) {
	client := pusher.Client{
		AppId:   "123123", // dummy Pusher AppID, provide real one to run this
		Key:     "xyz123", // dummy Pusher Key, provide real one to run this
		Secret:  "abc123", // dummy Pusher Secret, provide real one to run this
		Cluster: "eu",     // change when appropriate
		Secure:  true,
	}

	//  For every item in the list of commands
	for _, item := range commands {
		// send the item to pusher
		client.Trigger("my-channel", "my-event", convertData(item))
	}
}
