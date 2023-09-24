package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func RegisterCacheMiddleware(expiration time.Duration, cacheControl bool, maxBytes uint, methodsAllowed []string) (func(*fiber.Ctx) error) {
	return cache.New(cache.Config{
		Expiration: expiration,
		CacheControl: cacheControl,
		MaxBytes: maxBytes,
		Methods: methodsAllowed,
	})
}