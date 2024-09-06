package global

import (
	"net"
)

type User struct {
	Conn           net.Conn
	ConnClosed     bool
	Name           string
	MessageChannel chan Message
	Buffer         []byte
}

func NewUser(conn net.Conn) *User {
	u := User{
		Conn:           conn,
		Name:           conn.RemoteAddr().String(),
		MessageChannel: make(chan Message),
		Buffer:         make([]byte, 1024),
	}
	return &u
}

func (u *User) ResetBuffer() {
	u.Buffer = make([]byte, 1024)
}
