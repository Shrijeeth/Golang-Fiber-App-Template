package mail_provider

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/template_engine"
	"net/smtp"
	"os"
)

type SMTPMailer struct {
	Auth     smtp.Auth
	HostAddr string
	EmailID  string
}

func NewSMTPMailerProvider() (*SMTPMailer, error) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL_ID"), os.Getenv("SMTP_EMAIL_APP_PASSWORD"), os.Getenv("SMTP_HOST"))

	return &SMTPMailer{
		Auth:     auth,
		HostAddr: os.Getenv("SMTP_HOST_ADDR"),
		EmailID:  os.Getenv("SMTP_EMAIL_ID"),
	}, nil
}

func (m *SMTPMailer) SendMailWithTemplate(subject string, templateID string, fromMailDetails types.MailDetails, toMailDetails []types.MailDetails, templateVariables map[string]string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "\n"
	content, err := template_engine.ParseHTMLTemplate(templateID, templateVariables)
	if err != nil {
		return err
	}

	content = []byte(subject + mime + "\n" + string(content))

	var toMail []string
	for _, recipient := range toMailDetails {
		toMail = append(toMail, recipient.Email)
	}

	err = smtp.SendMail(m.HostAddr, m.Auth, m.EmailID, toMail, content)
	if err != nil {
		return err
	}

	return nil
}
