package email

import (
	"bytes"
	"log"

	"github.com/go-mail/mail"
	"github.com/jaytaylor/html2text"
	"github.com/spf13/viper"
	"github.com/vanng822/go-premailer/premailer"
)

// Mailer is a SMTP mailer.
type Mailer struct {
	client *mail.Dialer
	from   Email
}

// NewMailer returns a configured SMTP Mailer.
func NewMailer() (*Mailer, error) {
	if err := parseTemplates(); err != nil {
		return nil, err
	}

	smtp := struct {
		Host     string
		Port     int
		User     string
		Password string
	}{
		viper.GetString("mail.smtp_host"),
		viper.GetInt("mail.smtp_port"),
		viper.GetString("mail.smtp_user"),
		viper.GetString("mail.smtp_password"),
	}

	s := &Mailer{
		client: mail.NewPlainDialer(smtp.Host, smtp.Port, smtp.User, smtp.Password),
		from:   NewEmail(viper.GetString("mail.from_name"), viper.GetString("mail.from_address")),
	}

	if smtp.Host == "" {
		log.Println("SMTP host not set => printing emails to stdout")
		debug = true
		return s, nil
	}

	d, err := s.client.Dial()
	if err == nil {
		d.Close()
		return s, nil
	}
	return nil, err
}

// Send sends the mail via smtp.
func (m *Mailer) Send(email *message) error {
	if debug {
		log.Println("To:", email.to.Address)
		log.Println("Subject:", email.subject)
		log.Println(email.text)
		return nil
	}

	msg := mail.NewMessage()
	msg.SetAddressHeader("From", email.from.Address, email.from.Name)
	msg.SetAddressHeader("To", email.to.Address, email.to.Name)
	msg.SetHeader("Subject", email.subject)
	msg.SetBody("text/plain", email.text)
	msg.AddAlternative("text/html", email.html)

	return m.client.DialAndSend(msg)
}

// message struct holds all parts of a specific email message.
type message struct {
	from     Email
	to       Email
	subject  string
	template string
	content  interface{}
	html     string
	text     string
}

// parse parses the corrsponding template and content
func (m *message) parse() error {
	buf := new(bytes.Buffer)
	if err := templates.ExecuteTemplate(buf, m.template, m.content); err != nil {
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
	m.html = html

	text, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	m.text = text
	return nil
}
