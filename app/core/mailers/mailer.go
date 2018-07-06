package mailers

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/sknv/next/app/core/initers"
)

type Mailer struct {
	From string
	Addr string
	Auth smtp.Auth
}

func NewMailer() *Mailer {
	cfg := initers.GetConfig()
	return &Mailer{
		From: cfg.MailFrom,
		Addr: cfg.GetMailAddr(),
		Auth: smtp.PlainAuth(
			"", cfg.MailUsername, cfg.MailPassword, cfg.MailHost,
		),
	}
}

func (m *Mailer) ExecuteTemplate(name string, data interface{}) []byte {
	bytes, err := initers.GetHTML().ExecuteTemplateToBytes(name, data)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (m *Mailer) Deliver(email *email.Email) {
	// Log an email for the development mode.
	if !initers.GetConfig().IsRelease() {
		log.Printf("[INFO] deliver email: %s to %s", email.Text, email.To)
		return
	}

	// Actually deliver an email.
	email.From = m.From
	if err := email.Send(m.Addr, m.Auth); err != nil {
		log.Print("[ERROR] deliver email: ", err)
	}
}
