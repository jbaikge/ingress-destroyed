package imap

import (
	"code.google.com/p/rsc/imap"
	"github.com/jbaikge/ingress-destroyed/message"
	"log"
)

type Imap struct {
	Client *imap.Client
	Inbox  *imap.Box
}

func Login(server, user, passwd string) (i *Imap, err error) {
	i = &Imap{}
	i.Client, err = imap.NewClient(imap.Unencrypted, server, user, passwd, "")
	if err != nil {
		return
	}
	i.Inbox = i.Client.Inbox()
	return
}

func (i *Imap) Messages(msgs chan *message.Message) {
	log.Println("Starting IMAP Message poll")
	inbox := i.Client.Inbox()
	if err := inbox.Check(); err != nil {
		log.Printf("Error checking inbox: %s", err)
	}
}
