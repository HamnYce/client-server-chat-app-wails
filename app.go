package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
)

// App struct
type App struct {
	ctx      context.Context
	conn     net.Conn
	io       bufio.ReadWriter
	messages []string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.conn = nil
	a.messages = make([]string, 0)
}

func (a *App) ConnectToChatroom() {
	if a.conn != nil {
		return
	}

	conn, err := net.Dial("tcp4", "localhost:9090")
	if err != nil {
		conn = nil
		return
	}
	a.conn = conn
	a.io = *bufio.NewReadWriter(bufio.NewReader(a.conn), bufio.NewWriter(a.conn))

	log.Println("Successfully connected to server")
}

func (a *App) DisconnectFromChatroom() {
	a.conn.Close()
	a.conn = nil
}

func (a *App) SendMsgToChatRoom(msg string) {
	a.io.WriteString(msg)
	a.io.Flush()
}

func (a *App) ListenForMessage() {
	if !a.IsConnected() {
		return
	}
	a.io.Reader.Reset(a.conn)
	str, err := a.io.ReadString(0)

	// _, err := a.conn.Read(a.buffer)
	if err == io.EOF {
		a.DisconnectFromChatroom()
	} else if err != nil {
		log.Println(err)
		a.DisconnectFromChatroom()
	}
	log.Println("Received From Server:", str)
	a.messages = append(a.messages, str)
}

func (a *App) GetMessages() []string {
	return a.messages
}

func (a *App) IsConnected() bool {
	return a.conn != nil
}
