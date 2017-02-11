package robot

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

// Robot is an instance of a Sphero robot
type Robot struct {
	Driver     *bb8.BB8Driver
	commands   Commands
	bleAdaptor *ble.ClientAdaptor
}

// New creates a new instance of the sphero robot and connects using the
// given name
func New(name string) *Robot {
	bleAdaptor := ble.NewClientAdaptor(name)
	driver := bb8.NewDriver(bleAdaptor)

	return &Robot{Driver: driver, bleAdaptor: bleAdaptor}
}

// Execute a list of commands
func (r *Robot) Execute(commands Commands) {
	work := func() {
		commands.Execute(r.Driver)
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{r.bleAdaptor},
		[]gobot.Device{r.Driver},
		work,
	)

	robot.Start()
}
