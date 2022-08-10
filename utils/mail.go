package utilities

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendMail(to string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Account Activation")
	mailer.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		port,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	return dialer.DialAndSend(mailer)
}
