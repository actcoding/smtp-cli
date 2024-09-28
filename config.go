package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host    string   `json:"host"`
	Port    uint16   `json:"port"`
	User    string   `json:"username"`
	Pass    string   `json:"password"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
}

var (
	defaultFile = "smtp-cli.json"

	defaultConfig = &Config{
		Port:    465,
		Subject: "Neue Anmeldung auf dem Server",
	}
)

func LoadConfig(file *string) (*Config, error) {
	var (
		fileToOpen = file
		result     = defaultConfig
	)

	if fileToOpen == nil {
		fileToOpen = &defaultFile
	}

	contents, err := os.ReadFile(*fileToOpen)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
