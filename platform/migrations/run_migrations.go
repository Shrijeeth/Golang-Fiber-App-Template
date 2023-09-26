package migrations

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/models"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
)

func RunMigrations() {
	configs.DbClient.AutoMigrate(&models.User{}) //nolint:errcheck
}
