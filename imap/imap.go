package imap

import (
	"code.google.com/p/rsc/imap"
	"fmt"
	"github.com/jbaikge/ingress-destroyed/config"
	"github.com/jbaikge/ingress-destroyed/message"
	"log"
	"time"
)

type Imap struct {
	Client  *imap.Client
	Inbox   *imap.Box
	oldMsgs map[*imap.Msg]bool
}

func Login(server, user, passwd string) (i *Imap, err error) {
	i = &Imap{
		oldMsgs: make(map[*imap.Msg]bool),
	}
	i.Client, err = imap.NewClient(imap.Unencrypted, server, user, passwd, "")
	if err != nil {
		return
	}
	i.Inbox = i.Client.Inbox()
	return
}

func (c *Imap) LiveMessages(msgChan chan *imap.Msg) {
	log.Println("Starting IMAP Message poll")
	for {
		log.Print("Polling...")

		<-time.After(config.Imap.Refresh)
	}
}

func (c *Imap) AllMessages() (msgs []*imap.Msg, err error) {
	log.Println("Checking for new messages...")
	if err = c.Inbox.Check(); err != nil {
		err = fmt.Errorf("Error checking inbox: %s", err)
		return
	}
	msgs = c.Inbox.Msgs()
	log.Printf("Inbox contains %d messages", len(msgs))
	return
}

func (c *Imap) NewMessages() (msgs []*imap.Msg, err error) {
	all, err := c.AllMessages()
	if err != nil {
		return
	}
	for _, m := range all {
		if _, ok := c.oldMsgs[m]; ok {
			m.Flags |= imap.FlagDeleted
		} else {
			msgs = append(msgs, m)
		}
		c.oldMsgs[m] = false
	}
	for i := range c.oldMsgs {
		if c.oldMsgs[i] {
			delete(c.oldMsgs, i)
		} else {
			c.oldMsgs[i] = true
		}
	}
	log.Printf("%d new message(s)", len(msgs))
	return
}

func ToMessage(msg *imap.Msg) (m *message.Message) {
	m = &message.Message{
		Id:   fmt.Sprint(msg.UID),
		Date: msg.Date,
	}
	for _, c := range msg.Root.Child {
		switch c.Type {
		case "text/html":
			m.HTML = c.Text()
		case "text/plain":
			m.Text = c.Text()
		}
	}
	return
}
