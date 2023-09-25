package authentication

import (
	validator "github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg"
	"github.com/gofiber/fiber/v2"
)

type authentication struct{}

type AuthenticationControllerInterface interface {
	SignUp(c *fiber.Ctx) error
}

func GetAuthenticationServiceInstance() AuthenticationControllerInterface {
	return new(authentication)
}

func (auth *authentication) SignUp(c *fiber.Ctx) error {
	request := &SignUpDto{}
	if err := validator.ParseAndValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": true,
			"error":   err.Error(),
		})
	}

	return nil
}
