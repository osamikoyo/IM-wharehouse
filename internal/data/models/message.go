package models

import "gopkg.in/gomail.v2"

type Message struct{
	*gomail.Message
}

type Msg struct{
	From string
	To string
	CC []string
	Subject string
	Body string
}

func NewMessage(from, to string, addressHeader []string, subject string, body string) *Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", addressHeader[0], addressHeader[1])
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return &Message{m}
}