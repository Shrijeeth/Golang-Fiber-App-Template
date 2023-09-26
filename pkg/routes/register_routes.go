package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	router := app.Group("/api/v1")

	RegisterAuthenticateRoute(router)
	NotFoundRoute(app)
}
