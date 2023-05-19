package main

import "log"

type From struct {
	Name string
	Mail string
}

type Mail struct {
	From    From
	To      []string
	Subject string
	Body    string
}

func newFrom(name string, mail string) *From {
	return &From{
		Name: name,
		Mail: mail,
	}
}

func newMail(from From, to []string, suject string, body string) *Mail {
	return &Mail{
		From:    from,
		To:      to,
		Subject: suject,
		Body:    body,
	}
}

func (m *Mail) Send() {
	log.Println(m.From.Name, m.From.Mail)
	log.Println("send email")
}

func main() {
	from := newFrom("name", "from@address")
	mail := newMail(*from, []string{"to@address"}, "subject", "body")
	mail.Send()
}
