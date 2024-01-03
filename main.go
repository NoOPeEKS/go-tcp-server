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
	buffer := make([]byte, 64)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			log.Printf("ERROR: COULD NOT READ FROM CONNECTION: %s", err)
		}
		//channel <- buffer[0:n]
		//x := <-channel
		_, error := conn.Write(buffer[0:n])
		if error != nil {
			conn.Close()
			log.Printf("ERROR: COULD NOT READ FROM CONNECTION: %s", error)
		}
	}
}

func main() {
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatalf("ERROR: COULD NOT LISTEN TO PORT: %s", err)
	}

	//message := make(chan []byte)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: COULD NOT ACCEPT CONNECTION %s", conn.RemoteAddr().String())
		}

		log.Printf("ACCEPTED CONNECTION FROM %s", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}
