package migrations

import (
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/app/src/models"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/configs"
)

func RunMigrations() {
	configs.DbClient.AutoMigrate(&models.User{}) //nolint:errcheck
}
