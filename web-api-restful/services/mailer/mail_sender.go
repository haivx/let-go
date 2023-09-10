package mailer

import (
	config "final-project/config"
	"fmt"
)

type Information struct {
	Username string
	Password string
}

func SendGMail(info Information) {
	config, _ := config.LoadConfig(".")
	InitGmail(&SMTPConfig{
		Host:     config.MAIL_HOST,
		From:     config.MAIL_SENDER_USERNAME,
		Password: config.MAIL_SENDER_PASSWORD,
		Port:     config.MAIL_PORT,
	})

	data := map[string]interface{}{
		"title":    "mr/ms",
		"customer": info.Username,
		"info":     []Information{info},
	}

	err := Emailer.SendHTMLEmail([]string{config.MAIL_TO}, config.MAIL_SUBJECT, data, "template.html")
	if err != nil {
		fmt.Println("Could not send email: ", err)
	}
}
