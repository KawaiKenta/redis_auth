package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
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
	server.Username = config.Mail.UserAddress
	server.Password = config.Mail.Password
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = true
	server.ConnectTimeout = config.Mail.ConnectionTimeout
	server.SendTimeout = config.Mail.SendTimeout

	smtpClient, err = server.Connect()
	if err != nil {
		log.Fatalf("smtp client error: %v", err)
	}
}

func SendEmailVerifyMail(c *gin.Context, to string, token string) error {
	templateFileName := "views/verify_email.html"
	// TODO: ハードコーディング
	url := fmt.Sprintf("%s%s?uuid=%s", c.Request.Host, "/user/verify", token)
	data := struct {
		Token string
		Url   string
	}{Token: token, Url: url}

	subject := "メールアドレスを認証してください"
	if err := sendHtmlEmail(templateFileName, data, to, subject); err != nil {
		return err
	}
	return nil
}

// func SendResetPasswordMail(to string, token string) error {

// }

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
	from := fmt.Sprintf("%s <%s>", config.Mail.UserName, config.Mail.UserAddress)
	email.SetFrom(from).AddTo(to).SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	// send email
	if err := email.Send(smtpClient); err != nil {
		return err
	}
	return nil
}
