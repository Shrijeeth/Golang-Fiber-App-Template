package mail_provider

import (
	"context"
	"errors"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/mailersend/mailersend-go"
	"os"
	"time"
)

type MailerSend struct {
	Client *mailersend.Mailersend
}

func NewMailerSendProvider() (*MailerSend, error) {
	apiKey := os.Getenv("MAILERSEND_API_KEY")
	if apiKey == "" {
		return nil, errors.New("Invalid MailerSend Api Key")
	}

	client := mailersend.NewMailersend(apiKey)
	if client == nil {
		return nil, errors.New("Error Initializing MailerSend Client")
	}

	return &MailerSend{
		Client: client,
	}, nil
}

func (m *MailerSend) SendMailWithTemplate(subject string, templateID string, fromMailDetails types.MailDetails, toMailDetails []types.MailDetails, templateVariables map[string]string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sender := mailersend.From{
		Name:  fromMailDetails.Name,
		Email: fromMailDetails.Email,
	}

	recipients := []mailersend.Recipient{}
	for _, toMail := range toMailDetails {
		recipients = append(recipients, mailersend.Recipient{
			Name:  toMail.Name,
			Email: toMail.Email,
		})
	}

	templateSubstitutions := []mailersend.Substitution{}
	personalizationVariables := make(map[string]interface{})
	for key, value := range templateVariables {
		templateSubstitutions = append(templateSubstitutions, mailersend.Substitution{
			Var:   key,
			Value: value,
		})
		personalizationVariables[key] = value
	}
	variables := []mailersend.Variables{
		{
			Email:         fromMailDetails.Email,
			Substitutions: templateSubstitutions,
		},
	}

	personalization := []mailersend.Personalization{
		{
			Email: fromMailDetails.Email,
			Data:  personalizationVariables,
		},
	}

	message := m.Client.Email.NewMessage()
	message.SetFrom(sender)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)

	_, err := m.Client.Email.Send(ctx, message)

	return err
}
