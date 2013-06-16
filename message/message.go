package message

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
