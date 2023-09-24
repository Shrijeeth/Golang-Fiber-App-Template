package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func getMySqlConnectionString() string {
	host := os.Getenv("MYSQL_DB_HOST")
	user := os.Getenv("MYSQL_DB_USER")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	port := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, dbName)
	return connectionString
}

func MySqlConnect() (*gorm.DB, error) {
	connectionString := getMySqlConnectionString()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return db, err
}
