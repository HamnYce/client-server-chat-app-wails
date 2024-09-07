package global

import (
	"bytes"
	"fmt"
)

type Message struct {
	Sender  string
	Message string
	Time    string
}

const (
	delim string = "\u1000"
)

func (m Message) ToByteSlice() []byte {
	return ([]byte(fmt.Sprint(m.Sender + delim + m.Message + delim + m.Time)))
}

func NewMessageFromByteSlice(b []byte) (m Message) {
	chunks := bytes.Split(b, []byte(delim))
	return Message{
		Sender:  string(chunks[0]),
		Message: string(chunks[1]),
		Time:    string(chunks[2]),
	}
}
