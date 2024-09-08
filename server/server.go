package main

import (
	g "chatapp/global"
	"log"
)

var cr = g.NewChatroom()

/*
* main Flow
1) init message_history.json
2) init logs.txt
3) bind to external tailscale host
4) and begin listening for connections
4) pass connections into goroutine handler
*/
func main() {
  log.Println("Server is listening on "+g.HOST+":"+g.PORT)
	for {
		conn, err := cr.Server.Accept()
		if err != nil {
		}

		u := g.NewUser(conn)

		log.Printf("Successfully connected to %s\n", u.Name)

		cr.AddUser(u)
		go cr.MessageListener(u)
		go cr.MessageWriter(u)
	}
}
