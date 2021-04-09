package mailer

import (
	"log"

	"github.com/go-mail/mail"
	"github.com/spf13/viper"
)

// Mailer is a SMTP mailer.
type mailerImpl struct {
	client *mail.Dialer
}

// NewMailer returns a configured SMTP Mailer.
func NewMailer() Mailer {
	err := ParseTemplates()
	if err != nil {
		log.Fatal(err)
	}
	return &mailerImpl{
		client: mail.NewPlainDialer(
			viper.GetString("email.smtp_host"),
			viper.GetInt("email.smtp_port"),
			viper.GetString("email.smtp_user"),
			viper.GetString("email.smtp_password"),
		),
	}
}

// Send sends the mail via smtp.
func (m *mailerImpl) Send(email *MailboxEmail) error {
	msg := mail.NewMessage()
	msg.SetAddressHeader("From", email.From.Address, email.From.Name)
	msg.SetAddressHeader("To", email.To.Address, email.To.Name)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/plain", email.Text)
	msg.AddAlternative("text/html", email.HTML)

	return m.client.DialAndSend(msg)
}
