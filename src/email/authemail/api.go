package authemail

import (
	"time"

	"github.com/jobbox-tech/recruiter-api/email/mailer"
)

// AuthEmail interface
type AuthEmail interface {
	SendLoginEmail(to mailer.Recipient, content LoginEmail) error
}

// LoginEmail defines content for login token email template.
type LoginEmail struct {
	Email  string
	Name   string
	URL    string
	Token  string
	Expiry time.Time
}
