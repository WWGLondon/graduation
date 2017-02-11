package main

import "time"

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

	// create list of commands

	//r := robot.New("BB-7D60")
	//r.Execute(commands)

}
