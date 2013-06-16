package main

import (
	"flag"
	"github.com/jbaikge/ingress-destroyed/action"
	"github.com/jbaikge/ingress-destroyed/config"
	"github.com/jbaikge/ingress-destroyed/mail"
	"github.com/jbaikge/ingress-destroyed/mail/mbox"
	"github.com/jbaikge/ingress-destroyed/storage/csv"
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
	flag.Parse()

	if config.Mbox == "" {
		log.Fatal("No path for mbox file defined.")
	}

	f, err := os.Open(config.Mbox)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	msgChan := make(chan *mail.Message)
	box := &mbox.Mbox{f}
	log.Printf("Box ready. Parsing...")
	go box.Messages(msgChan)
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

func Store(a *action.Action) {
	for _, ch := range stores {
		ch <- a
	}
}
