package postgres

import (
	"github.com/jbaikge/ingress-destroyed/action"
	"github.com/jbaikge/ingress-destroyed/config"
	"testing"
	"time"
)

func init() {
	config.Init()
}

func TestPostgres(t *testing.T) {
	p, err := Open(config.Storage.Postgres)
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
}

func TestInsert(t *testing.T) {
	s, err := Open(config.Storage.Postgres)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	s.DB.Exec("DELETE FROM actions WHERE message_id = 1")

	a := &action.Action{
		Agent: "jbaikge",
		Enemy: "enemyNo1",
		Time:  time.Now(),
	}
	a.Id.MessageId, a.Id.ActionId = "1", 1
	if err := s.Save(a); err != nil {
		t.Fatal(err)
	}
}
