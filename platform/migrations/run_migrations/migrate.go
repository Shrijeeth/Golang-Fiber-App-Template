package main

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/migrations"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	// Load Database
	err = configs.InitDb()
	if err != nil {
		log.Panicf("Error initializing database: %s", err)
	}
	defer configs.CloseDb() //nolint:errcheck

	err = migrations.RunMigrations()
	if err != nil {
		log.Panicf("Error running migrations: %s", err)
	}
}
