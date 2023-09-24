package routes

import "github.com/gofiber/fiber/v2"

func RegisterAuthenticateRoute(router fiber.Router) {
	router.Group("/authenticate")
}
