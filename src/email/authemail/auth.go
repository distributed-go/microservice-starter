package authemail

import (
	"fmt"

	"github.com/jobbox-tech/recruiter-api/email/mailer"
	"github.com/spf13/viper"
)

type authEmail struct {
	mailer mailer.Mailer
}

// NewAuthEmail returns implementation for AuthEmail
func NewAuthEmail() AuthEmail {
	return &authEmail{
		mailer: mailer.NewMailer(),
	}
}

// SendLoginEmail creates and sends a login token email with provided template content.
func (m *authEmail) SendLoginEmail(to mailer.Recipient, content LoginEmail) error {
	msg := &mailer.MailboxEmail{
		From:     mailer.Recipient{Name: viper.GetString("email.from_name"), Address: viper.GetString("email.from_address")},
		To:       to,
		Subject:  fmt.Sprintf("Login To %s Recruitment Platform", viper.GetString("website.product_name")),
		Template: "loginToken",
		Content:  content,
	}

	if err := msg.Parse(); err != nil {
		return err
	}

	return m.mailer.Send(msg)
}
