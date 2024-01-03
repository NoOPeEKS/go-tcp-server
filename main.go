package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6969"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for connections in desired port
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatalf("ERROR: COULD NOT LISTEN TO PORT: %s", err)
	}

	log.Printf("Listening for connections on port %s", CONN_PORT)

	messages := make(chan Message)

	// Handle the message channel through the server
	go server(messages)

	for {
		// Accept every connection to the chatroom
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: COULD NOT ACCEPT CONNECTION %s", conn.RemoteAddr().String())
		}
		log.Printf("ACCEPTED CONNECTION FROM %s", conn.RemoteAddr().String())

		// Create a message, attach it the connection incoming and send it through the channel
		messages <- Message{
			Conn: conn,
		}

		// Handle the message channel through the client
		go client(conn, messages)
	}
}
