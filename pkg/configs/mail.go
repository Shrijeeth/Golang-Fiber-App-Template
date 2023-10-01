package configs

import (
	"errors"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/mail_provider"
	"os"
	"strconv"
)

type MailProvider interface {
	SendMailWithTemplate(subject string, templateID string, fromMailDetails types.MailDetails, toMailDetails []types.MailDetails, templateVariables map[string]string) error
}

var MailClient MailProvider

func InitMailClient() error {
	mailProvider := os.Getenv("MAIL_PROVIDER")
	var err error

	switch mailProvider {
	case "mailer_send":
		MailClient, err = mail_provider.NewMailerSendProvider()
	case "smtp":
		MailClient, err = mail_provider.NewSMTPMailerProvider()
	default:
		MailClient, err = nil, errors.New("Invalid Mail Provider")
	}

	return err
}

func CloseMailClient() error {
	mailProvider := os.Getenv("MAIL_PROVIDER")
	var err error

	switch mailProvider {
	case "mailer_send":
		err = nil
	default:
		err = errors.New("Invalid Mail Provider")
	}

	return err
}

func IsMailClientRequired() bool {
	isMailClientRequired, _ := strconv.Atoi(os.Getenv("MAIL_CLIENT_REQUIRED"))
	return isMailClientRequired == 1
}
