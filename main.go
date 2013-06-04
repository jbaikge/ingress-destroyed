package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Point struct {
	Agent     []byte
	Destroyer []byte
	Time      time.Time
	Count     int
	Location  Location
}

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
		msgPoint := &Point{
			Agent:     ExtractName(msg.HTML),
			Destroyer: ExtractDestroyer(msg.HTML),
		}
		msgPoint.Time, _ = msg.Message.Header.Date()

		locs, err := URLLocations(ExtractLinks(msg.HTML))
		if err != nil {
			log.Fatal(err)
		}
		for _, l := range locs {
			p := &Point{}
			*p = *msgPoint
			p.Location = *l
			//fmt.Printf("[%d.%d] %s %s %s <%0.6f,%0.6f>\n", msgCount, i, d.Format(time.Stamp), string(name), string(destroyer), l.Lat, l.Lon)
			fmt.Println(p)
		}

		msgCount++
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("\t\t{location: new google.maps.LatLng(%0.6f, %0.6f), weight: %d},", p.Location.Lat, p.Location.Lon, 1)
}
