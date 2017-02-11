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

// c1 := robot.Command{ElapsedTime: Milliseconds(1000), Do: func() { driver.Roll(Walking, West) }}
// c2 := robot.Command{ElapsedTime: Milliseconds(1000), Do: func() { driver.Roll(Running, South) }}
// c3 := robot.Command{ElapsedTime: Milliseconds(1000), Do: func() { driver.Roll(Crawling, East) }}
// c := robot.Commands{c1, c2, c3}
func getCommands(driver *bb8.BB8Driver) robot.Commands {
	c := robot.Commands{}

	c.Do(func() {
		fmt.Println("someing")
		driver.Roll(Walking, West)
	}).For(Milliseconds(3000))

	c.Do(func() {
		fmt.Println("else")
		driver.Roll(Running, South)
	}).For(Milliseconds(3000))

	c.Do(func() {
		fmt.Println("here")
		driver.Roll(Walking, East)
	}).For(Milliseconds(3000))

	c.Do(func() {
		fmt.Println("Stop")
		driver.Stop()
	}).For(Milliseconds(100))

	return c
}
