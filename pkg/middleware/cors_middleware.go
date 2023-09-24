package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func RegisterCORSMiddleware(origins []string, methods []string) fiber.Handler {
	return cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     strings.Join(origins, ", "),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     strings.Join(methods, ","),
	})
}
