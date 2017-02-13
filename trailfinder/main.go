// Team 1 solution from https://github.com/MarckK/wwggraduation/blob/master/main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

var commands Commands

type Command struct {
	Direction string
	Mode      string
	Order     int
}

type Commands []Command

func (a Commands) Len() int {
	return len(a)
}

func (a Commands) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Commands) Less(i, j int) bool {
	return a[i].Order < a[j].Order
}

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/WWGLondon/graduation/solution/map/release_party_map.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	commands = Commands{}
	err = json.Unmarshal(data, &commands)

	sort.Sort(commands)

	http.HandleFunc("/senddata", sendData)

	http.ListenAndServe(":9000", http.DefaultServeMux)

}

func sendData(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(commands)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	instructions := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://localhost:9000/senddata", instructions)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)

}
