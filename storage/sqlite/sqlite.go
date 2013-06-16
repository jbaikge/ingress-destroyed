package sqlite

import (
	"database/sql"
	"github.com/jbaikge/ingress-destroyed/action"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

type SQLite struct {
	DB *sql.DB
}

var create = `CREATE TABLE IF NOT EXISTS actions (
	message_id   TEXT,
	action_id    INTEGER,
	agent        TEXT,
	enemy        TEXT,
	time         TEXT,
	damage_type  TEXT,
	damage_count INTEGER,
	point_lat    REAL,
	point_lon    REAL,
	endpoint_lat REAL,
	endpoint_lon REAL,
	PRIMARY KEY(message_id, action_id)
)`

var insert = `INSERT OR IGNORE INTO actions VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

func Listener(filename string, wg *sync.WaitGroup) (ch chan *action.Action, err error) {
	s, err := Open(filename)
	if err != nil {
		return
	}
	ch = make(chan *action.Action)
	go func(s *SQLite, ch chan *action.Action) {
		wg.Add(1)
		for a := range ch {
			if err := s.Save(a); err != nil {
				log.Print(err)
			}
		}
		s.Close()
		wg.Done()
	}(s, ch)
	return
}

func Open(filename string) (s *SQLite, err error) {
	s = &SQLite{}
	if s.DB, err = sql.Open("sqlite3", filename); err != nil {
		return
	}
	_, err = s.DB.Exec(create)
	return
}

func (s *SQLite) Close() {
	s.DB.Close()
}

func (s *SQLite) Save(a *action.Action) (err error) {
	_, err = s.DB.Exec(insert,
		a.Id.MessageId,
		a.Id.ActionId,
		a.Agent,
		a.Enemy,
		a.Time.Format("2006-01-02 15:04:05.999"),
		a.Damage.Type.String(),
		a.Damage.Count,
		a.Point.Lat, a.Point.Lon,
		a.EndPoint.Lat,
		a.EndPoint.Lon,
	)
	return
}
