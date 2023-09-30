package routes

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthenticateRoute(router fiber.Router) {
	authentication := controllers.GetAuthenticationServiceInstance()

	router = router.Group("/authenticate")
	router.Post("/signup", authentication.SignUp)
	router.Post("/signin", authentication.SignIn)

	router = router.Group("/google")
	router.Get("/signin", authentication.GoogleSignIn)
	router.Get("/redirect", authentication.GoogleCallback)
}
