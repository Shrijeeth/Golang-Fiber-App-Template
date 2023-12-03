package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getMySqlConnectionString(isTest bool) string {
	var host, user, password, port, dbName string
	if isTest {
		host = os.Getenv("MYSQL_TEST_DB_HOST")
		user = os.Getenv("MYSQL_TEST_DB_USER")
		password = os.Getenv("MYSQL_TEST_DB_PASSWORD")
		port = os.Getenv("MYSQL_TEST_DB_PORT")
		dbName = os.Getenv("MYSQL_TEST_DB_NAME")
	} else {
		host = os.Getenv("MYSQL_DB_HOST")
		user = os.Getenv("MYSQL_DB_USER")
		password = os.Getenv("MYSQL_DB_PASSWORD")
		port = os.Getenv("MYSQL_DB_PORT")
		dbName = os.Getenv("MYSQL_DB_NAME")
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, dbName)
	return connectionString
}

func MySqlConnect(isTest bool) (*gorm.DB, error) {
	connectionString := getMySqlConnectionString(isTest)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return db, err
}
