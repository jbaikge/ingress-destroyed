package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
	"time"
)

type imap struct {
	Host      string
	Password  string
	Username  string
	DeleteOld bool
	Once      bool
	Refresh   time.Duration
}

type storage struct {
	SQLite string
	CSV    string
}

var (
	Imap    = imap{}
	Mbox    = ""
	Storage = storage{}
	UseIMAP = false
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := ReadFromFile(os.Getenv("HOME") + "/.config/ingress-destroyed.yaml"); err != nil {
		log.Fatalf("Error processing configuration: %s", err)
	}

	if Imap.Refresh == 0 {
		Imap.Refresh = 5 * time.Minute
	}

	flag.DurationVar(&Imap.Refresh, "imap.refresh", Imap.Refresh, "Duration between refresh intervals")
	flag.BoolVar(&Imap.DeleteOld, "imap.deleteold", Imap.DeleteOld, "Delete old (processed) messages")
	flag.BoolVar(&UseIMAP, "useimap", UseIMAP, "true = Use IMAP config; false = Use Mbox config")
	flag.StringVar(&Imap.Host, "imap.host", Imap.Host, "IMAP Host")
	flag.StringVar(&Imap.Username, "imap.username", Imap.Username, "IMAP Username")
	flag.StringVar(&Imap.Password, "imap.password", Imap.Password, "IMAP Password")
	flag.BoolVar(&Imap.Once, "imap.once", Imap.Once, "Only pull IMAP messages once, do not poll")
	flag.StringVar(&Mbox, "mbox", Mbox, "User mbox path (usually /var/mail/user)")
	flag.StringVar(&Storage.CSV, "storage.csv", Storage.CSV, "CSV Path")
	flag.StringVar(&Storage.SQLite, "storage.sqlite", Storage.SQLite, "SQLite Database Path")
}

func ReadFromFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Error reading %s: %s", filename, err)
	}

	c := struct {
		Imap    *imap
		Mbox    string
		Storage *storage
		UseIMAP bool
	}{
		Imap:    &Imap,
		Storage: &Storage,
	}

	if err := goyaml.Unmarshal(b, &c); err != nil {
		return fmt.Errorf("Error processing %s: %s", filename, err)
	}
	Mbox = c.Mbox
	UseIMAP = c.UseIMAP
	return
}
