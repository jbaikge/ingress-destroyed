package mail

import (
	"time"
)

type Message struct {
	Id   string
	Date time.Time
	From string
	Text []byte
	HTML []byte
}

func (m Message) FromNiantic() bool {
	return m.From == "Niantic Project Operations <ingress-support@google.com>"
}
