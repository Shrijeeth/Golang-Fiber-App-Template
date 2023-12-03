package main

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/routes"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/jobs"
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
	err = configs.InitDb(false)
	if err != nil {
		log.Panicf("Error initializing database: %s", err)
	}
	defer configs.CloseDb() //nolint:errcheck

	// Load Cache Database
	isCacheRequired := configs.IsCacheRequired()
	if isCacheRequired {
		err = configs.InitRedisClient()
		if err != nil {
			log.Panicf("Error initializing cache: %s", err)
		}
		defer configs.CloseRedisClient() //nolint:errcheck
	}

	// Load Cloud Storage Object
	isCloudStorageObjectRequired := configs.IsCloudStorageObjectRequired()
	if isCloudStorageObjectRequired {
		err = configs.InitCloudObjectStorage()
		if err != nil {
			log.Panicf("Error initializing cloud storage object: %s", err)
		}
		defer configs.CloseCloudObjectStorage() //nolint:errcheck
	}

	// Load Mail Client
	isMailClientRequired := configs.IsMailClientRequired()
	if isMailClientRequired {
		err = configs.InitMailClient()
		if err != nil {
			log.Panicf("Error initializing email client: %s", err)
		}
		defer configs.CloseMailClient() //nolint:errcheck
	}

	// Load Job Worker
	isJobWorkerRequired := jobs.IsJobWorkerRequired()
	if isJobWorkerRequired {
		err := jobs.InitWorkerPool()
		if err != nil {
			log.Panicf("Error initializing job worker: %s", err)
		}
		defer jobs.CloseWorkerPool() //nolint:errcheck
	}

	// Initialize fiber app
	app := fiber.New()

	//Register App Routes
	routes.RegisterRoutes(app)

	// Start the Server
	app.Listen(os.Getenv("APP_ADDR")) //nolint:errcheck
}
