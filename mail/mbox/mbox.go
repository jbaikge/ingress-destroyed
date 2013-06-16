package mbox

import (
	message "github.com/jbaikge/ingress-destroyed/mail"
	"io"
)

type Mbox struct {
	Reader io.Reader
}

func (m *Mbox) Messages(msgChan chan *message.Message) {
	blockChan := make(chan []byte)
	Blocks(m.Reader, blockChan)
	for block := range blockChan {
		msg := &message.Message{}
		toMessage(block, msg)
		msgChan <- msg
	}
}
