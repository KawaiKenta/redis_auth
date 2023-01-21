package mail

import (
	"bytes"
	"html/template"
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
	"kk-rschian.com/redis_auth/config"
)

var (
	smtpClient *mail.SMTPClient
	err        error
)

func SetUp() {
	server := mail.NewSMTPClient()
	// SMTP Server
	server.Host = config.Mail.Host
	server.Port = config.Mail.Port
	server.Username = config.Mail.User
	server.Password = config.Mail.Password
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = config.Mail.ConnectionTimeout
	server.SendTimeout = config.Mail.SendTimeout

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	// TODO: これを無効化する必要があるかも

	smtpClient, err = server.Connect()

	if err != nil {
		log.Fatalf("smtp client error: %v", err)
	}
}

func SendEmailVerifyMail(to string, token string) error {
	templateFileName := "views/verify_email.html"
	data := struct{ Token string }{Token: token}
	subject := "メールアドレスを認証してください"
	if err := sendHtmlEmail(templateFileName, data, to, subject); err != nil {
		return err
	}
	return nil
}

func sendHtmlEmail(templateFileName string, data any, to string, subject string) error {
	// get html from template
	template, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := template.Execute(buf, data); err != nil {
		return err
	}
	body := buf.String()

	// create email
	email := mail.NewMSG()
	email.SetFrom(config.Mail.User).AddTo(to).SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	// send email
	if err := email.Send(smtpClient); err != nil {
		return err
	}
	return nil
}
