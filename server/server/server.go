package main

import (
	g "chat-server/global"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

var (
	users = make(map[string]*g.User)
)

func broadcastMessage(msg g.Message) {
	for _, u := range users {
		u.MessageChannel <- msg
	}
}
func DisconnectUser(u *g.User) {
	u.ConnClosed = true
	u.Conn.Close()
	close(u.MessageChannel)
	delete(users, u.Name)
	log.Printf("Successfully Disconnected from %s\n", u.Name)
}

/*
* main Flow
1) init message_history.json
2) init logs.txt
3) bind to external tailscale host
4) and begin listening for connections
4) pass connections into goroutine handler
*/
func main() {
	server, err := net.Listen("tcp4", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
		}

		u := g.NewUser(conn)

		log.Printf("Successfully connected to %s\n", u.Name)

		users[u.Name] = u
		go chatroomListener(u)
		go chatroomWriter(u)
	}
}

/*
* Listener Handler Flow
1) initiate buffer that will exist for the duration of the connection
2) listen for msgs sent by connection in the form of name:msg
3) save those messages into message_history and save on disc
*/
func chatroomListener(u *g.User) {
	for {
		u.ResetBuffer()
		_, err := u.Conn.Read(u.Buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("error taking in input from, ", u.Conn.LocalAddr().String())
			break
		}

		msg := g.Message{
			Sender:  u.Name,
			Message: string(u.Buffer),
			Time:    time.Now().Format("DateTime"),
		}

		log.Printf("Listener: Received msg %s from %s", msg.Message, u.Name)

		if strings.ToLower(msg.Message) == "exit now" {
			break
		}

		broadcastMessage(msg)
	}
	DisconnectUser(u)
}

/*
* Writer Handler Flow
1) initiate message counter for user to keep track of how far they are along the conversation
2) constantly check if the length of the chat history is greater than there message counter
3) if message_history length is greater then send all messaages until they are caught up (including there own)
4) else keep looping we can also create a list of channels and send to those
*/
func chatroomWriter(u *g.User) {
	for msg := range u.MessageChannel {
		if u.ConnClosed {
			break
		}
		log.Printf("Writer: %+v\n", msg)
		u.Conn.Write([]byte(msg.Message))
	}
}
