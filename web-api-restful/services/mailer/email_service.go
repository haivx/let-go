package mailer

import (
	"fmt"
	"net/smtp"
)

type MailSender interface {
	SendHTMLEmail(to []string, subject string, data map[string]interface{}, template string) error
}

var Emailer MailSender

type SMTPConfig struct {
	Host     string
	From     string
	Password string
	Port     int
}

type GmailSTMP struct {
	config *SMTPConfig
}

func InitGmail(config *SMTPConfig) {
	Emailer = GmailSTMP{
		config: config,
	}
}

func (gmail GmailSTMP) SendHTMLEmail(to []string, subject string, data map[string]interface{}, template string) error {
	emailAuth := smtp.PlainAuth("", gmail.config.From, gmail.config.Password, gmail.config.Host)
	body, err := renderHTML(data, template)
	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", gmail.config.Host, gmail.config.Port)

	if err := smtp.SendMail(addr, emailAuth, gmail.config.From, to, msg); err != nil {
		return err
	}

	return nil
}
