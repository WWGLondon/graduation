package pusher

import (
	"encoding/json"
	"fmt"
	"net/url"

	"golang.org/x/net/websocket"
)

var listeners chan *Connection

// init is a package level function to start the run loop
func init() {
	listeners = make(chan *Connection)
	go runLoop()
}

// runLoop is an internal function to keep connection listeners alive
func runLoop() {
	for newListener := range listeners {
		go newListener.listen()
	}
}

// Connection is responsible for connecting to pusher and handling the receipt
// of messages
type Connection struct {
	ws            *websocket.Conn
	subscriptions map[string]chan InMessage
}

// Connect to the pusher server
func (c *Connection) Connect(url *url.URL, origin string) error {
	c.subscriptions = make(map[string]chan InMessage)

	ws, err := websocket.Dial(url.String(), "", origin)
	if err != nil {
		return err
	}

	c.ws = ws

	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		return err
	}
	fmt.Printf("Received: %s.\n", msg[:n])

	return nil
}

// Subscribe subscribes to a channel, when a message is received there will be
// a new message pushed to the channel
func (c *Connection) Subscribe(channel string) (chan InMessage, error) {
	subscribeChannel := make(chan InMessage, 100) // use a buffered channel to avoid blocking

	event := OutMessage{
		Event: SubscribeEvent,
		Data: Data{
			Channel: channel,
		},
	}

	data, err := json.Marshal(event)
	if err != nil {
		return subscribeChannel, err
	}

	_, err = c.ws.Write(data)
	if err != nil {
		return subscribeChannel, err
	}

	c.subscriptions[channel] = subscribeChannel

	go func(conn *Connection) {
		listeners <- conn
	}(c)

	return subscribeChannel, nil
}

// listen is an internal method which is called by the run loop
func (c *Connection) listen() {
	var buffer = make([]byte, 512)

	for {
		if n, err := c.ws.Read(buffer); err == nil {
			if n > 1 {
				event := InMessage{}
				err := json.Unmarshal(buffer[:n], &event)
				if err != nil {
					fmt.Println(err)
				}

				subscriber, ok := c.subscriptions[event.Channel]
				if ok {
					subscriber <- event
				}

			}
		}
	}
}
