package routes

import (
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/app/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthenticateRoute(router fiber.Router) {
	router = router.Group("/authenticate")
	authentication := controllers.GetAuthenticationServiceInstance()

	router.Post("/signup", authentication.SignUp)
	router.Post("/signin", authentication.SignIn)
}
