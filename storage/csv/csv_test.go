package csv

import (
	"github.com/jbaikge/ingress-destroyed/action"
	"os"
	"testing"
	"time"
)

func TestCSV(t *testing.T) {
	filename := "/tmp/test.csv"
	os.Remove(filename)

	c, err := Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	a := &action.Action{
		Agent: "jbaikge",
		Enemy: "EnemyNo1",
		Time:  time.Now(),
	}
	if err := c.Save(a); err != nil {
		t.Fatal(err)
	}
}
