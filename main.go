package main

import (
	"flag"
	"log"
	"os"
)

var (
	mboxFilename = flag.String("f", "/var/mail/jake", "Path to mbox file to parse")
)

func main() {
	flag.Parse()

	f, err := os.Open(*mboxFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	msgs := make(chan []byte)
	go mboxMessageBlocks(f, msgs)
	for msg := range msgs {
		log.Print(string(msg))
	}
}
