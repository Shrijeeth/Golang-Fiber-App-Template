package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(
		RegisterHelmetMiddleware(helmet.Config{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
		}),
		RegisterCORSMiddleware([]string{
			"http://127.0.0.1:3000",
		}, []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}),
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}),
		RegisterLimiterMiddleware(25, 30*time.Second),
		RegisterCacheMiddleware(10*time.Second, true, 0, []string{
			fiber.MethodGet,
		}),
		RegisterBlackListingMiddleware("../../blacklist.txt"),
		RegisterWhiteListingMiddleware("../../whitelist.txt"),
	)
}
