package utilities

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendMail(to string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("MAIL_SENDER_NAME"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Account Activation")
	mailer.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("MAIL_SMTP_HOST"),
		port,
		os.Getenv("MAIL_AUTH_EMAIL"),
		os.Getenv("MAIL_AUTH_PASSWORD"),
	)

	return dialer.DialAndSend(mailer)
}
