package main

import (
	"flag"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Data struct {
	Host       string
	User       string
	RemoteUser string
	RemoteHost string
	Tty        string
	Timestamp  time.Time
}

var flagConfig string
var flagTemplate string

func main() {
	flag.StringVar(&flagConfig, "config", "smtp-cli.json", "The config file to use.")
	flag.StringVar(&flagTemplate, "template", "template.gotmpl", "The config file to use.")
	flag.Parse()

	config, err := LoadConfig(&flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(log.DebugLevel)

	data := getData()
	SendMail(config, flagTemplate, data)
}

func getData() Data {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return Data{
		Host:       hostname,
		User:       os.Getenv("PAM_USER"),
		RemoteUser: os.Getenv("PAM_RUSER"),
		RemoteHost: os.Getenv("PAM_RHOST"),
		Tty:        os.Getenv("PAM_TTY"),
		Timestamp:  time.Now(),
	}
}
