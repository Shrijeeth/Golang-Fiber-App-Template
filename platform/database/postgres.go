package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getPostgresConnectionString() string {
	host := os.Getenv("POSTGRES_DB_HOST")
	user := os.Getenv("POSTGRES_DB_USER")
	password := os.Getenv("POSTGRES_DB_PASSWORD")
	port := os.Getenv("POSTGRES_DB_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	sslMode := os.Getenv("POSTGRES_DB_SSL_MODE")
	if sslMode != "disable" {
		sslMode += " sslrootcert=" + os.Getenv("POSTGRES_DB_SSL_CERTIFICATE")
	}
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata", host, user, password, dbName, port, sslMode)
	return connectionString
}

func PostgresConnect() (*gorm.DB, error) {
	connectionString := getPostgresConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	return db, err
}
