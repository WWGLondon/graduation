package pusher

import (
	"encoding/json"
	"fmt"
	"net/url"

	"golang.org/x/net/websocket"
)

var listeners chan *Connection

func init() {
	listeners = make(chan *Connection)
	go runLoop()
}

func runLoop() {
	for newListener := range listeners {
		fmt.Println("New Listener")
		go newListener.listen()
	}
}

type Connection struct {
	ws            *websocket.Conn
	subscriptions map[string]chan Message
}

func (c *Connection) Connect(url *url.URL, origin string) error {
	c.subscriptions = make(map[string]chan Message)

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

func (c *Connection) Subscribe(channel string) (chan Message, error) {
	subscribeChannel := make(chan Message)

	event := Message{
		Event: SubscribeEvent,
		Data: Data{
			Channel: channel,
		},
	}

	data, err := json.Marshal(event)
	if err != nil {
		return subscribeChannel, err
	}

	fmt.Printf("%#v\n", string(data))

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

func (c *Connection) listen() {
	fmt.Println("Start listen")
	var buffer = make([]byte, 512)

	for {
		if n, err := c.ws.Read(buffer); err == nil {
			if n > 1 {
				fmt.Printf("Received: %s.\n", buffer[:n])
			}
		}
	}
}
