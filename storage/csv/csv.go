package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/jbaikge/ingress-destroyed/action"
	"log"
	"os"
)

type CSV struct {
	File   *os.File
	Writer *csv.Writer
}

const openFlags = os.O_RDWR | os.O_APPEND | os.O_CREATE

var fields = []string{
	"agent",
	"enemy",
	"datetime",
	"week",
	"weekday",
	"damage_type",
	"damage_count",
	"point_lat",
	"point_lon",
	"endpoint_lat",
	"endpoint_lon",
}

func Listener(filename string) (ch chan *action.Action, err error) {
	c, err := Open(filename)
	if err != nil {
		return
	}
	ch = make(chan *action.Action)
	go func(c *CSV, ch chan *action.Action) {
		for a := range ch {
			if err := c.Save(a); err != nil {
				log.Print(err)
			}
		}
		c.Close()
	}(c, ch)
	return
}

func Open(filename string) (f *CSV, err error) {
	f = &CSV{}
	if f.File, err = os.OpenFile(filename, openFlags, 0666); err != nil {
		return
	}
	f.Writer = csv.NewWriter(f.File)
	stat, err := f.File.Stat()
	if err != nil {
		return
	}
	if stat.Size() == 0 {
		f.Writer.Write(fields)
	}
	return
}

func (c *CSV) Close() {
	c.Writer.Flush()
	c.File.Close()
}

func (c *CSV) Save(a *action.Action) (err error) {
	_, week := a.Time.ISOWeek()
	record := []string{
		a.Agent,
		a.Enemy,
		a.Time.String(),
		fmt.Sprint(week),
		a.Time.Weekday().String(),
		a.Damage.Type.String(),
		fmt.Sprint(a.Damage.Count),
		fmt.Sprintf("%f", a.Point.Lat),
		fmt.Sprintf("%f", a.Point.Lon),
		fmt.Sprintf("%f", a.EndPoint.Lat),
		fmt.Sprintf("%f", a.EndPoint.Lon),
	}
	return c.Writer.Write(record)
}
