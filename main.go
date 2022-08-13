package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/bagasfathoni/go-email/sender/model"
)

type EmailConfig struct {
	Cred model.EmailCred
	Host string
	Port string
}

const EMAIL_TEMPLATE = `Subject: {{ .Subject }}
{{ .Body }}`

var t *template.Template

func main() {
	t = template.New("email")
	t.Parse(EMAIL_TEMPLATE)
	message := model.EmailMessage{
		From:    "address@mail.com",
		To:      []string{"destination@mail.com"},
		Subject: "Test",
		Body:    "This is a test",
	}
	var body bytes.Buffer
	t.Execute(&body, message)

	// Email config
	emailConfig := InitEmailConfig()
	auth := smtp.PlainAuth("", emailConfig.Cred.Address, emailConfig.Cred.Password, emailConfig.Host)

	// Send an email
	fmt.Printf("Sending an email from %s\n", emailConfig.Cred.Address)
	err := smtp.SendMail(emailConfig.Host+":"+emailConfig.Port, auth, message.From, message.To, body.Bytes())
	if err != nil {
		log.Fatalln(err)
	}
}

func InitEmailConfig() EmailConfig {
	cfg := new(EmailConfig)
	cfg.Cred.Address = os.Getenv("EMAIL_ADDRESS")
	cfg.Cred.Password = os.Getenv("EMAIL_PASSWORD")
	cfg.Host = os.Getenv("EMAIL_HOST")
	cfg.Port = os.Getenv("EMAIL_PORT")
	return *cfg
}
