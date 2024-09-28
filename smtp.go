package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func SendMail(config *Config, templateFile string, templateData any) error {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Debugf("Dial %s \n", address)

	connection, err := tls.Dial("tcp", address, &tls.Config{
		ServerName: config.Host,
	})
	if err != nil {
		return err
	}
	defer connection.Close()

	log.Debugf("Connect %s \n", address)
	client, err := smtp.NewClient(connection, config.Host)
	if err != nil {
		return err
	}
	defer client.Quit()

	log.Debug("Auth")
	auth := smtp.PlainAuth("", config.User, config.Pass, config.Host)
	err = client.Auth(auth)
	if err != nil {
		return err
	}

	log.Debug("Template")
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	for i, rcpt := range config.To {
		log.Debugf("Mail %d \n", i)
		log.Debugf("  From %s \n", config.From)
		err = client.Mail(config.From)
		if err != nil {
			return err
		}

		log.Debugf("  To %s \n", rcpt)
		err = client.Rcpt(rcpt)
		if err != nil {
			return err
		}

		log.Debug("  Send")
		w, err := client.Data()
		if err != nil {
			return err
		}

		log.Debug("  Headers")
		headers := make(map[string]string)
		headers["From"] = config.From
		headers["To"] = rcpt
		headers["Subject"] = config.Subject
		headers["Content-Type"] = "text/plain; charset=\"UTF-8\""
		for k, v := range headers {
			_, err = w.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
			if err != nil {
				return err
			}
		}

		log.Debug("  Body")
		err = tmpl.Execute(w, templateData)
		if err != nil {
			return err
		}

		err = w.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
