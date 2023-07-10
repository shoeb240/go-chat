# Basic Chat Room in Golang

Learn how to build a basic chat room in Golang with our easy-to-follow tutorial. This step-by-step video will guide you through the process of implementing client-server communication. No prior experience is needed. Let's get coding!

## Video Tutorial
This Chat Room is described step-by-step in a Youtube video tutorial [Basic Chat Room in Golang: Step by Step Tutorial || Part-1](https://www.youtube.com/watch?v=GPhlaMBAFrk)

## Chat Room Architecture
![Chat Room Architecture](https://github.com/shoeb240/go-chat/blob/main/assets/architecture.png)

In this architecture, the "client" represents an individual user connected to the chatroom.
It contains a WebSocket connection (conn) to communicate with the client, a channel (messageChan) to send and receive messages to and from the client, and a reference to the chatroom (room) the client belongs to.

room" represents the chatroom itself and holds information about the clients currently present in the room.
It maintains a map of clients (clients) and has three channels: joinChan for adding new clients to the room, broadcastChan for broadcasting messages to all clients in the room, and leaveChan for removing clients from the room.

The room handles chatroom operations. It listens to the channels and performs appropriate actions based on the received signals, such as adding new clients to the room, broadcasting messages to all clients, and removing clients from the room.

The client continuously reads incoming messages from the client's WebSocket connection and broadcasts them to all clients in the room using the broadcastChan.

The client continuously writes outgoing messages from the messageChan to the client's WebSocket connection.

This architecture allows multiple clients to join the chatroom, exchange messages with each other, and leave the chatroom seamlessly using WebSocket connections and goroutines.
