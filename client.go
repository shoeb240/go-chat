package main

import "github.com/gorilla/websocket"

type client struct {
	socket      *websocket.Conn
	messageChan chan []byte
	room        *room
}

func (c *client) readFromSocket() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		c.room.broadcastChan <- msg
	}
}

func (c *client) writeToSocket() {
	defer c.socket.Close()

	for msg := range c.messageChan {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}

}
