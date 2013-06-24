package config

import (
	"bytes"
	"testing"
	"time"
)

var contents = []byte(`
imap:
  host: mail.example.com
  password: pass
  username: usrnm
  deleteold: true
  once: true
  refresh: 10m
mbox: /var/mail/user
storage:
  sqlite: "/tmp/actions.sqlite"
  csv: "/tmp/actions.csv"
  postgres:
    host: localhost
    username: IngressDestroyed
    password: secret
    database: IngressDestroyed
useimap: true
`)

func TestConfig(t *testing.T) {
	r := bytes.NewReader(contents)
	if err := ReadConfig(r); err != nil {
		t.Fatal(err)
	}
	checks := []struct{ Got, Expect interface{} }{
		{Imap.Host, "mail.example.com"},
		{Imap.Password, "pass"},
		{Imap.Username, "usrnm"},
		{Imap.Once, true},
		{Imap.DeleteOld, true},
		{Imap.Refresh, 10 * time.Minute},
		{Mbox, "/var/mail/user"},
		{Storage.SQLite, "/tmp/actions.sqlite"},
		{Storage.CSV, "/tmp/actions.csv"},
		{Storage.Postgres.Host, "localhost"},
		{Storage.Postgres.Username, "IngressDestroyed"},
		{Storage.Postgres.Password, "secret"},
		{Storage.Postgres.Database, "IngressDestroyed"},
		{UseIMAP, true},
	}
	for i, c := range checks {
		if c.Got != c.Expect {
			t.Errorf("[%d] Failed. Got %v; Expected %v", i, c.Got, c.Expect)
		}
	}
}
