package main

import (
	"fmt"
	"time"

	"github.com/WWGLondon/graduation/gobot/robot"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

const (
	Crawling = 10
	Walking  = 20
	Running  = 30
)

const (
	North = 0
	East  = 90
	South = 180
	West  = 270
)

func Milliseconds(d int) time.Duration {
	return time.Duration(d) * time.Millisecond
}

//GODEBUG=cgocheck=0 go run main.go
func main() {
	bleAdaptor := ble.NewClientAdaptor("SK-AA49")
	driver := bb8.NewDriver(bleAdaptor)

	// TODO: Where would the data came from?
	//directions := []int{180, 60, 270, 180, 90}
	//step := 0

	command := getCommands(driver)

	work := func() {
		command.Execute(driver)
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{driver},
		work,
	)

	robot.Start()
}

func getCommands(driver *bb8.BB8Driver) robot.Commands {
	c := robot.Commands{}
	c.Do(func() {
		fmt.Println("Something")
	}).For(Milliseconds(1000))

	return c
}
