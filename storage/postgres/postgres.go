package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jbaikge/ingress-destroyed/action"
	"github.com/jbaikge/ingress-destroyed/config"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Postgres struct {
	DB *sql.DB
}

var create = `CREATE TABLE actions (
	message_id   BIGINT NOT NULL,
	action_id    SMALLINT NOT NULL,
	agent        VARCHAR(64),
	enemy        VARCHAR(64),
	time         TIMESTAMP WITH TIME ZONE,
	damage_type  VARCHAR(16) NOT NULL,
	damage_count SMALLINT NOT NULL,
	point        POINT NOT NULL,
	endpoint     POINT,
	PRIMARY KEY(message_id, action_id)
)`

var insert = `INSERT INTO actions VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

func Listener(cfg config.Postgres, wg *sync.WaitGroup) (ch chan *action.Action, err error) {
	p, err := Open(cfg)
	if err != nil {
		return
	}
	ch = make(chan *action.Action)
	go func(p *Postgres, ch chan *action.Action) {
		wg.Add(1)
		for a := range ch {
			if err := p.Save(a); err != nil {
				log.Print(err)
			}
		}
		p.Close()
		wg.Done()
	}(p, ch)
	return
}

func Open(cfg config.Postgres) (p *Postgres, err error) {
	log.Printf("%+v", cfg)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Database)

	p = &Postgres{}
	if p.DB, err = sql.Open("postgres", dsn); err != nil {
		return
	}
	// Super dangerous to not bubble the error here, but Postgres 8.3.23 does
	// not support IF NOT EXISTS
	p.DB.Exec(create)
	return
}

func (p *Postgres) Close() {
	p.DB.Close()
}

func (p *Postgres) Save(a *action.Action) (err error) {
	_, err = p.DB.Exec(insert,
		a.Id.MessageId,
		a.Id.ActionId,
		a.Agent,
		a.Enemy,
		a.Time.Format("2006-01-02 15:04:05.999-07:00"),
		a.Damage.Type.String(),
		a.Damage.Count,
		fmt.Sprintf("(%f, %f)", a.Point.Lat, a.Point.Lon),
		fmt.Sprintf("(%f, %f)", a.EndPoint.Lat, a.EndPoint.Lon),
	)
	return
}
