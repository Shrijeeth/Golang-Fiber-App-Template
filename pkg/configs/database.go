package configs

import (
	"errors"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/database"
	"gorm.io/gorm"
	"os"
)

var DbClient *gorm.DB

func InitDb(isTest bool) error {
	dbType := os.Getenv("DB_TYPE")
	var err error

	switch dbType {
	case "postgres":
		DbClient, err = database.PostgresConnect(isTest)
	case "mysql":
		DbClient, err = database.MySqlConnect(isTest)
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
