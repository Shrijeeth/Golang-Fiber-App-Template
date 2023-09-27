package migrations

import (
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/models"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
)

func RunMigrations() error {
	err := configs.DbClient.AutoMigrate(&models.User{}) //nolint:errcheck
	if err != nil {
		return err
	}

	fmt.Println("Migrations Completed Successfully")
	return nil
}
