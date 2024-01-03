package main

func server(messages chan Message) {
	// Create a map of clients to know who is in chat
	clients := map[string]*Client{}
	for {
		// Get the messages incoming from the channel
		msg := <-messages

		// Register the client who sent the message
		clients[msg.Conn.RemoteAddr().String()] = &Client{
			Conn: msg.Conn,
		}

		// Send the message to the other clients who did not send it so they can see it
		for _, client := range clients {
			if client.Conn.RemoteAddr().String() != msg.Conn.RemoteAddr().String() {
				client.Conn.Write([]byte(msg.Text))
			}
		}
	}
}
