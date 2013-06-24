package main

import (
	"flag"
	"fmt"
	"github.com/jbaikge/ingress-destroyed/action"
	"github.com/jbaikge/ingress-destroyed/config"
	"github.com/jbaikge/ingress-destroyed/mail"
	"github.com/jbaikge/ingress-destroyed/mail/imap"
	"github.com/jbaikge/ingress-destroyed/mail/mbox"
	"github.com/jbaikge/ingress-destroyed/storage/csv"
	"github.com/jbaikge/ingress-destroyed/storage/postgres"
	"github.com/jbaikge/ingress-destroyed/storage/sqlite"
	"log"
	"os"
	"sync"
)

var (
	stores = make([]chan *action.Action, 0, 2)
	wg     sync.WaitGroup
)

func main() {
	config.Init()
	flag.Parse()

	SetupStores()

	msgChan := make(chan *mail.Message)
	if config.UseIMAP {
		if err := SetupIMAP(msgChan); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := SetupMbox(msgChan); err != nil {
			log.Fatal(err)
		}
	}

	for msg := range msgChan {
		if !msg.FromNiantic() {
			continue
		}
		log.Printf("Parsing message: %s From %s", msg.Id, msg.From)
		actions, err := action.FromMessage(msg)
		if err != nil {
			log.Print(err)
			continue
		}
		for i, a := range actions {
			a.Id.MessageId = msg.Id
			a.Id.ActionId = i
			Store(a)
		}
	}
	CloseStores()
	wg.Wait()
}

func CloseStores() {
	for _, ch := range stores {
		close(ch)
	}
}

func SetupIMAP(msgChan chan *mail.Message) (err error) {
	if config.Imap.Host == "" {
		return fmt.Errorf("No IMAP host specified")
	}

	client, err := imap.Login(config.Imap.Host, config.Imap.Username, config.Imap.Password)
	if err != nil {
		return
	}

	go func() {
		log.Printf("IMAP ready. Parsing...")
		client.Messages(msgChan, config.Imap.DeleteOld)
		client.Client.Close()
	}()
	return
}

func SetupMbox(msgChan chan *mail.Message) (err error) {
	if config.Mbox == "" {
		return fmt.Errorf("No path for mbox file defined")
	}

	f, err := os.Open(config.Mbox)
	if err != nil {
		return
	}

	box := &mbox.Mbox{f}
	go func() {
		log.Printf("Mbox ready. Parsing...")
		box.Messages(msgChan)
		f.Close()
	}()
	return
}

func SetupStores() {
	log.Print("Setting up storage engines...")
	if path := config.Storage.CSV; path != "" {
		log.Print("Using CSV storage")
		ch, err := csv.Listener(path, &wg)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, ch)
		log.Print("CSV storage ready")
	}

	if path := config.Storage.SQLite; path != "" {
		log.Print("Using SQLite storage")
		ch, err := sqlite.Listener(path, &wg)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, ch)
		log.Print("SQLite storage ready")
	}

	if config.Storage.Postgres.Host != "" {
		log.Print("Using Postgres Storage")
		ch, err := postgres.Listener(config.Storage.Postgres, &wg)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, ch)
		log.Print("Postgres storage ready")
	}
}

func Store(a *action.Action) {
	for _, ch := range stores {
		ch <- a
	}
}
