package global

import (
	"net"
)

type User struct {
	Conn               net.Conn
	Name               string
	MessageChannel     chan Message
	LatestMessageIndex int
	Buffer             []byte
}

func NewUser(conn net.Conn) *User {
	u := User{
		Conn:               conn,
		Name:               conn.RemoteAddr().String(),
		MessageChannel:     make(chan Message),
		LatestMessageIndex: 0,
		Buffer:             make([]byte, 1024),
	}
	return &u
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) UnLoad() {
	close(u.MessageChannel)
	u.Conn.Close()
}
