package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

var (
	//smtpServer = "smtp.gmail.com"
	//smtpPort = "587"
	//user = "sei1rou@gmail.com"
	//pass = "conan@009"
	smtpServer = "smtp.shoeikai.com"
	smtpPort   = "465"
	user       = "stcheck@shoeikai.com"
	pass       = "65w$5TdG"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", user, pass, smtpServer)

	from := user
	to := "watanabe@shoeikai.com"
	title := "メール送信テスト"

	body := "メール届いたかな"

	header := make(map[string]string)
	header["Form"] = from
	header["To"] = to
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":"+smtpPort,
		auth,
		user,
		[]string{to},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}
