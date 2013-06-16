package sqlite

import (
	"database/sql"
	"github.com/jbaikge/ingress-destroyed/action"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	DB *sql.DB
}

var create = `CREATE TABLE IF NOT EXISTS actions (
	agent        TEXT,
	enemy        TEXT,
	time         TEXT,
	damage_type  TEXT,
	damage_count INTEGER,
	point_lat    REAL,
	point_lon    REAL,
	endpoint_lat REAL,
	endpoint_lon REAL
)`

var insert = `INSERT INTO actions VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
