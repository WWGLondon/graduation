package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WWGLondon/graduation/gobot/robot"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

const (
	// Crawling sets the speed to very slow
	Crawling = 20

	// Walking is normal speed
	Walking = 50

	// Running is a fast speed
	Running = 100
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

//GODEBUG=cgocheck=0 go run main.go
func main() {
	bleAdaptor := ble.NewClientAdaptor("SK-AA49")
	// bleAdaptor := ble.NewClientAdaptor("BB-7D60")
	driver := bb8.NewDriver(bleAdaptor)
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
		log.Println("Walking North")
		driver.Roll(Walking, North)
	}).For(Milliseconds(3000))

	c.Do(func() {
		log.Println("Crawling West")
		driver.Roll(Crawling, West)
	}).For(Milliseconds(3000))

	c.Do(func() {
		log.Println("Running North")
		driver.Roll(Running, North)
	}).For(Milliseconds(3000))

	c.Do(func() {
		log.Println("Running East")
		driver.Roll(Running, East)
	}).For(Milliseconds(3000))

	c.Do(func() {
		log.Println("Crawling South")
		driver.Roll(Crawling, South)
	}).For(Milliseconds(3000))

	c.Do(func() {
		log.Println("Walking North")
		driver.Roll(Walking, North)
	}).For(Milliseconds(3000))

	c.Do(func() {
		fmt.Println("Stop")
		driver.Stop()
	}).For(Milliseconds(100))

	return c
}

// # Team 4
// 1. Consruct a Sphero client
// 2. Decrypt directions -- maybe
// 3. Convert the data from pusher into Sphero commands
// 4. Deliver the code to the release party

// CRYPTO
//NORTH
//[char number][char number]...*[mesage id]
// 12142016*2
// divide by messge number, break into 2 number blocks, convert to letter, combine
