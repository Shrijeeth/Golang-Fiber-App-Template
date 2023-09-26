package main

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/routes"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	// Load Cache Database
	err = configs.InitRedisClient()
	if err != nil {
		log.Panicf("Error initializing cache: %s", err)
	}
	defer configs.CloseRedisClient() //nolint:errcheck

	// Run Database Migrations
	migrations.RunMigrations()

	// Initialize fiber app
	app := fiber.New()

	//Register App Routes
	routes.RegisterRoutes(app)

	// Start the Server
	app.Listen(os.Getenv("APP_ADDR")) //nolint:errcheck
}
