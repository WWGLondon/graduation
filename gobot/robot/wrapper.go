package robot

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

type Command struct {
	ElapsedTime time.Duration
	Do          func()
}

// Commands is a collection of Command
type Commands []Command

// GetCommand returns a command which is appropriate for the
// current elapsed time.
func (c Commands) GetCommand(elapsedTime time.Duration) Command {
	for _, command := range c {
		if elapsedTime <= command.ElapsedTime {
			return command
		}
	}

	return Command{}
}

// Do sets the things to do
func (c *Commands) Do(function func()) *Commands {
	*c = append(*c, Command{Do: function})
	return c
}

// For sets the duration to do things for
func (c *Commands) For(duration time.Duration) *Commands {
	latest := c.last(0)
	previous := c.last(1)

	latest.ElapsedTime = previous.ElapsedTime + duration

	return c
}

func (c Commands) last(index int) *Command {
	if len(c)-index > 0 {
		fmt.Println(len(c))
		return &c[len(c)-index-1]
	}

	return &Command{0, nil}
}

func (c Commands) Len() int           { return len(c) }
func (c Commands) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Commands) Less(i, j int) bool { return c[i].ElapsedTime < c[j].ElapsedTime }

// Execute tells your adventurer to executes the list of commands
func (c *Commands) Execute(driver *bb8.BB8Driver) {
	elapsed := 0 * time.Millisecond

	out := make(chan struct{})

	ticker := gobot.Every(50*time.Millisecond, func() {
		command := c.GetCommand(elapsed)
		if command.Do == nil {
			out <- struct{}{}
			return
		}

		command.Do()
		elapsed += 50 * time.Millisecond
		fmt.Println("Elapsed Time: ", elapsed)
	})

	<-out
	ticker.Stop()
}

// FlashLEDs flashes the LEDs for the robot
func FlashLEDs(driver *bb8.BB8Driver) {
	r := uint8(gobot.Rand(255))
	g := uint8(gobot.Rand(255))
	b := uint8(gobot.Rand(255))

	driver.SetRGB(r, g, b)
}
