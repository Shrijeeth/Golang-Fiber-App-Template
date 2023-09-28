package migrations

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/models"
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(
		"1695921509_create_users_table",
		func(db *gorm.DB) error {
			return db.AutoMigrate(&models.User{})
		},
		func(db *gorm.DB) error {
			return db.Migrator().DropTable(&models.User{})
		},
		MigrationUp,
	)
}
