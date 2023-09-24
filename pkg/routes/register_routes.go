package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	NotFoundRoute(app)
	router := app.Group("/api/v1")

	RegisterAuthenticateRoute(router)
}
