package main

import (
	"bufio"
	g "chatapp/global"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// taking in username
	// setting username at the very beginning to avoid userna
	fmt.Println("Enter your username")
	var username string
	fmt.Scan(&username)

	conn, err := net.Dial("tcp4", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("SET_NAME:" + username))

	log.Println("Successfully connected to server, you may begin to send messages")

	scanner := bufio.NewScanner(os.Stdin)
	go ListeningHandler(conn)
	// TODO: issue with receiving multiple messages
	// to recreate try joining server sending message then join again, for some reason it isnt workin
	// another test is to join with 2 clients at the same time and seeing what happens
	for {
		success := scanner.Scan()

		if !success || scanner.Err() == io.EOF {
			log.Println("Successfuly closed scanner")
			break
		} else if scanner.Err() != nil {
			log.Println(err)
			break
		}

		if len(scanner.Bytes()) == 0 {
			log.Println("Empty line received. If you would like to stop the program please press Ctrl+D")
			continue
		}

		// fmt.Println("Read:", scanner.Text())

		conn.Write(scanner.Bytes())
	}
}

func ListeningHandler(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		// listens and prints to stdio the messages from the server in parellel with sending
		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		msg := g.NewMessageFromByteSlice(buffer[:n])
		fmt.Printf("%s/%s: %s\n", strings.Split(msg.Time, " ")[1], msg.Sender, msg.Message)
	}
}
