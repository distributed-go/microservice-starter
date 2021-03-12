package mailer

import (
	"bytes"
	"html/template"

	"github.com/jaytaylor/html2text"
	"github.com/vanng822/go-premailer/premailer"
)

// Recipient holds the recipient email and name
type Recipient struct {
	Name    string
	Address string
}

// MailboxEmail holds the mailbox message to be sent
type MailboxEmail struct {
	From     Recipient
	To       Recipient
	Subject  string
	Template string
	Content  interface{}
	HTML     string
	Text     string
}

// Mailer ...
type Mailer interface {
	Send(email *MailboxEmail) error
}

// Parse parses the corresponding template and content
func (m *MailboxEmail) Parse() error {
	buf := new(bytes.Buffer)
	if err := templates.ExecuteTemplate(buf, m.Template, m.Content); err != nil {
		return err
	}
	prem, err := premailer.NewPremailerFromString(buf.String(), premailer.NewOptions())
	if err != nil {
		return err
	}

	html, err := prem.Transform()
	if err != nil {
		return err
	}
	m.HTML = html

	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	m.Text = text
	return nil
}

var (
	templates *template.Template
)
