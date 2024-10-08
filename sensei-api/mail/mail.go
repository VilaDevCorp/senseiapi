package mail

import (
	"net/smtp"
	"sensei/conf"
)

func SendMail(to string, subject string, message string) error {
	conf := conf.Get()

	auth := smtp.PlainAuth("", conf.Mail.User, conf.Mail.Pass, conf.Mail.SmtpHost)
	toArray := []string{to}
	msg := []byte("From: senseitrainingapp@gmail.com\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		message + "\r\n")

	err := smtp.SendMail(conf.Mail.SmtpHost+":"+conf.Mail.SmtpPort, auth, conf.Mail.User, toArray, msg)
	if err != nil {
		return err
	}
	return nil
}
