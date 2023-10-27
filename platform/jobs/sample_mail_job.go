package jobs

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/gocraft/work"
	"os"
)

func SampleMailJob(job *work.Job) error {
	if configs.IsMailClientRequired() {
		username := job.ArgString("username")
		email := job.ArgString("email")
		if err := job.ArgError(); err != nil {
			return err
		}

		from := types.MailDetails{
			Email: os.Getenv("APP_MAIL"),
			Name:  "Go-Fiber Template",
		}
		to := []types.MailDetails{
			{
				Name:  username,
				Email: email,
			},
		}
		templateVariables := map[string]string{
			"name": username,
		}
		err := configs.MailClient.SendMailWithTemplate("Registration Success", "sample-template.html", from, to, templateVariables)
		if err != nil {
			return err
		}
	}
	return nil
}
