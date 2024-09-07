package global

import (
	"encoding/csv"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	CHATROOM_PATH = "./chat.csv"
	HOST          = "localhost"
	PORT          = "9090"
)

type Chatroom struct {
	Server   net.Listener
	Users    map[string]*User
	Messages []Message
}

func NewChatroom() *Chatroom {
	server, err := net.Listen("tcp4", HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	return &Chatroom{
		Users:    make(map[string]*User),
		Messages: make([]Message, 0),
		Server:   server,
	}
}

// DATA

func (cr *Chatroom) AddUser(u *User) {
	cr.Users[u.Conn.RemoteAddr().String()] = u
}

func (cr *Chatroom) RemoveUser(u *User) {
	delete(cr.Users, u.Conn.RemoteAddr().String())
	u.UnLoad()
	log.Printf("Successfully Disconnected from %s\n", u.Name)
}

func (cr *Chatroom) AddMessage(msg Message) {
	cr.Messages = append(cr.Messages, msg)
}

// CONNECTION

func (cr *Chatroom) AcceptUsernameAndGreet(u *User) (success bool) {
	n, err := u.Conn.Read(u.Buffer)
	if err == io.EOF {
		return false
	} else if err != nil {
		log.Println("error taking in input from, ", u.Conn.LocalAddr().String())
		return false
	}

	_, name, _ := strings.Cut(string(u.Buffer[:n]), "SET_NAME:")
	log.Printf("Set name of %s to %s", u.Conn.RemoteAddr().String(), name)
	u.SetName(name)

	cr.AddMessage(NewMessage("Server", u.Name+" has joined the server. Welcome them!"))
	return true
}

func (cr *Chatroom) MessageListener(u *User) {
	if success := cr.AcceptUsernameAndGreet(u); success {
		cr.BroadcastNewMessages()
		for {
			n, err := u.Conn.Read(u.Buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Println("error taking in input from, ", u.Conn.LocalAddr().String())
				break
			}

			msg := NewMessage(u.Name, string(u.Buffer[:n]))

			log.Printf("Listener: Received msg %s from %s", msg.Message, u.Name)

			cr.AddMessage(msg)
			cr.BroadcastNewMessages()
		}
	}
	cr.RemoveUser(u)
}

func (cr *Chatroom) MessageWriter(u *User) {
	for msg := range u.MessageChannel {
		log.Printf("Sending %s to %s", msg.Message, u.Name)
		u.Conn.Write(msg.ToByteSlice())
	}
}

// UPDATING

func (cr *Chatroom) BroadcastNewMessages() {
	for _, u := range cr.Users {
		for _, msg := range cr.Messages[u.LatestMessageIndex:] {
			u.MessageChannel <- msg
			u.LatestMessageIndex++
		}
	}
}

// CSV RELATED FUNCTIONS

func NewChatroomFromCSV() *Chatroom {
	f, err := os.OpenFile(CHATROOM_PATH, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	cr := NewChatroom()
	cr.loadMessageHistoryFromRecords(records)
	return cr
}

func (cr *Chatroom) loadMessageHistoryFromRecords(records [][]string) {
	for _, record := range records {
		msg := Message{
			Sender:  record[0],
			Message: record[1],
			Time:    record[2],
		}
		cr.AddMessage(msg)
	}
}

func (cr Chatroom) ToCSV() {
	f, err := os.OpenFile(CHATROOM_PATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	records := cr.ToRecords()
	csvWriter := csv.NewWriter(f)
	csvWriter.WriteAll(records)
}

func (cr Chatroom) ToRecords() (records [][]string) {
	records = make([][]string, 0)
	for _, msg := range cr.Messages {
		records = append(records, msg.ToRecord())
	}

	return
}
