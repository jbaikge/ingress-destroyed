package main

import (
	"flag"
	"log"
	"os"
)

var (
	MsgStart     = []byte("From ")
	mboxFilename = flag.String("f", "/var/mail/jake", "Path to mbox file to parse")
)

func main() {
	flag.Parse()

	f, err := os.Open(*mboxFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	blocks := make(chan []byte)
	go mboxMessageBlocks(f, blocks)
	for block := range blocks {
		msg, err := toMessage(block)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s %s", msg.Message.Header.Date(), )
	}
}
