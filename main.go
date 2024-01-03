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

func handleConnection(conn net.Conn) {
	conn.Write([]byte("Hello friend!"))
	conn.Close()
}

func main() {
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatalf("ERROR: COULD NOT LISTEN TO PORT: %s", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: COULD NOT ACCEPT CONNECTION %s", conn.RemoteAddr().String())
		}

		go handleConnection(conn)
	}
}
