package configs

import (
	"errors"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/database"
	"gorm.io/gorm"
	"os"
)

var DbClient *gorm.DB

func InitDb() error {
	dbType := os.Getenv("DB_TYPE")
	var err error

	switch dbType {
	case "postgres":
		DbClient, err = database.PostgresConnect()
	case "mysql":
		DbClient, err = database.MySqlConnect()
	default:
		DbClient, err = nil, errors.New("Invalid DB Type")
	}

	return err
}

func CloseDb() error {
	if DbClient == nil {
		return nil
	}

	db, err := DbClient.DB()
	db.Close()

	return err
}
