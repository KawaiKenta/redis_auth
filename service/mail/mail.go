package utils

import "gopkg.in/gomail.v2"

func TestMail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "kawai@iseyama.online")
	msg.SetHeader("To", "Amagaki29870@gmail.com")
	msg.SetHeader("Subject", "テストメール")
	msg.SetBody("text/html", "<b>This is the body of the mail. 日本語はどうか</b>")

	n := gomail.NewDialer(
		"sv8682.xserver.jp",
		587,
		"no-reply@iseyama.online",
		"iseyama4336",
	)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}

// we should use go-smtp instead of gomail.v2 because this repo is frozen since 2016
