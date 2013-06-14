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
	t.Logf("Inbox name: %s", c.Inbox.Name)
}
