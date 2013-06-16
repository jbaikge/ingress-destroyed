package imap

import (
	"github.com/jbaikge/ingress-destroyed/config"
	"testing"
)

func TestLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("Short tests enabled")
	}
	c, err := Login(config.Imap.Host, config.Imap.Username, config.Imap.Password)
	if err != nil {
		t.Fatalf("Could not login: %s", err)
	}
	defer c.Client.Close()
	t.Logf("Inbox name: %s", c.Inbox.Name)
}

func TestAllMessages(t *testing.T) {
	if testing.Short() {
		t.Skip("Short tests enabled")
	}
	c, err := Login(config.Imap.Host, config.Imap.Username, config.Imap.Password)
	if err != nil {
		t.Fatalf("Could not login: %s", err)
	}
	defer c.Client.Close()
	msgs, err := c.AllMessages()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%d messages", len(msgs))
}

func TestNewMessages(t *testing.T) {
	if testing.Short() {
		t.Skip("Short tests enabled")
	}
	c, err := Login(config.Imap.Host, config.Imap.Username, config.Imap.Password)
	if err != nil {
		t.Fatalf("Could not login: %s", err)
	}
	defer c.Client.Close()
	msgs0, err := c.NewMessages()
	if err != nil {
		t.Fatal(err)
	}
	msgs1, err := c.NewMessages()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(msgs0); l == 0 {
		t.Fatalf("Got zero messages for first call")
	}
	if l := len(msgs1); l != 0 {
		t.Fatalf("Second call should be zero; got: %d", l)
	}
}

func TestSingleMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("Short tests enabled")
	}
	c, err := Login(config.Imap.Host, config.Imap.Username, config.Imap.Password)
	if err != nil {
		t.Fatalf("Could not login: %s", err)
	}
	defer c.Client.Close()
	msgs, err := c.AllMessages()
	if err != nil {
		t.Fatal(err)
	}
	m := ToMessage(msgs[0])
	t.Logf("ID: %s", m.Id)
	t.Logf("Date: %s", m.Date)
	t.Logf("From: %s", m.From)
	t.Logf("HTML len: %d", len(m.HTML))
	t.Logf("Text len %d", len(m.Text))
}
