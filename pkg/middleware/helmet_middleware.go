package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func RegisterHelmetMiddleware(helmetConfig helmet.Config) (func(c *fiber.Ctx) error) {
	return helmet.New(helmetConfig)
}