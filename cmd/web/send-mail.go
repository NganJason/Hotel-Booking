package main

import (
	"fmt"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/NganJason/hotel-booking/internal/models"
)

func listenForMail() {
	go func() {
		fmt.Println("Listening for email")
		for {
			m := <- app.MailChan
			sendMsg(m)
		}
	}()

}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextPlain, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent")
	}
}