package sqlite

import (
	"github.com/jbaikge/ingress-destroyed/action"
	"os"
	"testing"
	"time"
)

var filename = "/tmp/test.db"

func TestSqlite(t *testing.T) {
	os.Remove(filename)

	s, err := Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
}

func TestInsert(t *testing.T) {
	os.Remove(filename)

	s, err := Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	a := &action.Action{
		Agent: "jbaikge",
		Enemy: "enemyNo1",
		Time:  time.Now(),
	}
	if err := s.Save(a); err != nil {
		t.Fatal(err)
	}
}
