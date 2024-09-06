package main

import (
	"bufio"
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

		fmt.Println("Read:", scanner.Text())

		conn.Write(scanner.Bytes())

		_, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		fmt.Println("Server:", string(buffer))
	}
}
