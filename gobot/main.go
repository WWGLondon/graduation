// Team 4 solution from https://github.com/sangeetha28/graduation/blob/master/gobot/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/WWGLondon/graduation/gobot/robot"
)

const (
	// Crawling sets the speed to very slow
	Crawling = 20

	// Walking is normal speed
	Walking = 20

	// Running is a fast speed
	Running = 20
)

const (
	// North is a compass north
	North = 0

	// East is compass east
	East = 90

	// South is compass south
	South = 180

	// West is compas west
	West = 270
)

// Milliseconds is a convenience function which returns time.
func Milliseconds(d int) time.Duration {
	return time.Duration(d) * time.Millisecond
}

// var Commands []robot.Command
var robby *robot.Robot

//GODEBUG=cgocheck=0 go run main.go
func main() {
	robby = robot.New("SK-AA49") // Update to actual bot name (might need to use ble explorer to lookup the name)

	go newWebServer()

	timer := time.NewTimer(10000 * time.Millisecond)
	<-timer.C

	checkAndRun()
}

type Data struct {
	Direction string
	Mode      string
}

func getHeading(h string) uint16 {
	fmt.Println(h)
	switch h {
	case "North":
		return North
	case "East":
		return East
	case "South":
		return South
	case "West":
		return West
	default:
		return North
	}
}

func getSpeed(s string) uint8 {
	switch s {
	case "Running":
		return Running
	case "Walking":
		return Walking
	case "Crawling":
		return Crawling
	default:
		return Walking
	}
}

var input []Data

func checkAndRun() {

	if len(input) == 0 {
		return
	}

	commands := robot.Commands{}
	for _, d := range input {
		func(d Data) {
			commands.Do(func() {
				fmt.Println(d.Direction)
				s := getSpeed(d.Mode)
				h := getHeading(d.Direction)
				robby.Driver.Roll(s, h)
			}).For(2000 * time.Millisecond)
		}(d)
	}

	robby.Execute(commands)
}

func newWebServer() {
	http.HandleFunc("/input", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		input2 := Data{}
		err = json.Unmarshal(data, &input2)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Printf("%#v\n", input2)
		input = append(input, input2)
	})

	_ = http.ListenAndServe(":7001", http.DefaultServeMux)
	log.Fatal("Boom")
}
