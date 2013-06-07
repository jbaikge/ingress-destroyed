package config

import (
	"flag"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

type configType struct {
	Imap    login
	Mbox    string
	Storage storage
}

type login struct {
	Host     string
	Password string
	Username string
}

type storage struct {
	SQLite string
	CSV    string
	MySQL  login
}

var Config = configType{}

func init() {
	var configFile = os.Getenv("HOME") + "/.config/ingress-destroyed.yaml"
	if f, err := os.Open(configFile); err == nil {
		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal("Error reading", configFile, ":", err)
		}
		if err := goyaml.Unmarshal(b, &Config); err != nil {
			log.Fatal("Error processing", configFile, ":", err)
		}
	} else {
		log.Printf("Couldn't read %s: %s", configFile, err)
	}

	flag.StringVar(&Config.Imap.Host, "imap.host", Config.Imap.Host, "IMAP Host")
	flag.StringVar(&Config.Imap.Username, "imap.username", Config.Imap.Username, "IMAP Username")
	flag.StringVar(&Config.Imap.Password, "imap.password", Config.Imap.Password, "IMAP Password")
	flag.StringVar(&Config.Mbox, "mbox", Config.Mbox, "User mbox path (usually /var/mail/user)")
	flag.StringVar(&Config.Storage.MySQL.Host, "storage.mysql.host", Config.Storage.MySQL.Host, "MySQL Host")
	flag.StringVar(&Config.Storage.MySQL.Username, "storage.mysql.username", Config.Storage.MySQL.Username, "MySQL Username")
	flag.StringVar(&Config.Storage.MySQL.Password, "storage.mysql.password", Config.Storage.MySQL.Password, "MySQL Password")
	flag.StringVar(&Config.Storage.CSV, "storage.csv", Config.Storage.CSV, "CSV Path")
	flag.StringVar(&Config.Storage.SQLite, "storage.sqlite", Config.Storage.SQLite, "SQLite Database Path")
}
