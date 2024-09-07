package main

import (
	"bufio"
	g "chatapp/global"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Successfully connected to server")

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		fmt.Print("User Input: ")
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

		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		msg := g.NewMessageFromByteSlice(buffer[:n])
		fmt.Printf("Sender:%s, Message:%s, Time:%s\n", msg.Message, msg.Sender, msg.Time)
	}
}
