package authentication

import (
	"log"
	"testing"

	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/platform/migrations"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthenticationService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Service Test Suite")
}

var _ = BeforeEach(func() {
	err := godotenv.Load("../../../../../.env.example")
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	err = configs.InitDb(true)
	if err != nil {
		log.Panicf("Error loading test db: %s", err)
	}

	err = migrations.RunMigrations()
	if err != nil {
		log.Panicf("Error running test migrations: %s", err)
	}
})

var _ = AfterEach(func() {
	migrations.RollbackMigrations() //nolint:errcheck

	configs.CloseDb() //nolint:errcheck
})