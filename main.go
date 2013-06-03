package main

import (
	"flag"
	"fmt"
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
	msgCount := 0
	go mboxMessageBlocks(f, blocks)
	for block := range blocks {
		msg, err := toMessage(block)
		if err != nil {
			log.Fatal(err)
		}
		d, _ := msg.Message.Header.Date()
		name := ExtractName(msg.HTML)
		destroyer := []byte(`unknown`)
		locs, err := URLLocations(ExtractLinks(msg.HTML))
		if err != nil {
			log.Fatal(err)
		}
		for i, l := range locs {
			fmt.Printf("[%d.%d] %s %s %s <%0.6f,%0.6f>\n", msgCount, i, d, string(name), string(destroyer), l.Lat, l.Lon)
		}
		msgCount++
	}
}
