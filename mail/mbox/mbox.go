package mbox

import (
	"github.com/jbaikge/ingress-destroyed/mail"
	"io"
)

type Mbox struct {
	Reader io.Reader
}

func (m *Mbox) Messages(msgChan chan *mail.Message) {
	blockChan := make(chan []byte)
	Blocks(m.Reader, blockChan)
	for block := range blockChan {
		msg := &mail.Message{}
		toMessage(block, msg)
		msgChan <- msg
	}
}
