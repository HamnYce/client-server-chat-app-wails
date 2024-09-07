package global

import (
	"bytes"
	"fmt"
	"time"
)

type Message struct {
	Sender  string
	Message string
	Time    string
}

const (
	delim string = "\u1000"
)

func NewMessage(sender, message string) Message {
	return Message{
		Sender:  sender,
		Message: message,
		Time:    time.Now().Format(time.DateTime),
	}
}

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

func (m Message) ToRecord() []string {
	return []string{
		m.Sender,
		m.Message,
		m.Time,
	}
}
