package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	clients       map[*client]bool
	joinChan      chan *client
	leaveChan     chan *client
	broadcastChan chan []byte
}

func newRoom() *room {
	return &room{
		clients:       make(map[*client]bool),
		joinChan:      make(chan *client),
		leaveChan:     make(chan *client),
		broadcastChan: make(chan []byte),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.joinChan:
			r.clients[client] = true

		case client := <-r.leaveChan:
			delete(r.clients, client)
			close(client.messageChan)

		case msg := <-r.broadcastChan:
			for client := range r.clients {
				client.messageChan <- msg
			}
		}

	}

}

const (
	socketBufferSize  = 1024
	messageBufferSize = 1024
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &client{
		socket:      socket,
		messageChan: make(chan []byte, messageBufferSize),
		room:        r,
	}

	r.joinChan <- client

	defer func() { r.leaveChan <- client }()

	go client.writeToSocket()

	client.readFromSocket()

}
