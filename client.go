package main

import (
	"net"
)

type Message struct {
	Conn net.Conn
	Text string
}

type Client struct {
	Conn net.Conn
}

func client(conn net.Conn, messages chan Message) {
	// Create a buffer to store the messages the client sends
	buffer := make([]byte, 64)
	for {
		// Read the messages into the buffer
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			return
		}

		// Convert the messages from bytes to a string
		text := string(buffer[0:n])

		// Attach the string to a message struct and send it through the channel
		messages <- Message{
			Conn: conn,
			Text: text,
		}
	}
}
