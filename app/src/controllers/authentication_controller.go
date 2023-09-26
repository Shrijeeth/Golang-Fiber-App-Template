package controllers

import (
	"context"
	authenticationService "github.com/Shrijeeth/Personal-Finance-Tracker-App/app/src/services/authentication"
	validator "github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg"
	"github.com/gofiber/fiber/v2"
)

type authentication struct{}

type AuthenticationControllerInterface interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}

func GetAuthenticationServiceInstance() AuthenticationControllerInterface {
	return new(authentication)
}

func (auth *authentication) SignUp(c *fiber.Ctx) error {
	request := &authenticationService.SignUpDto{}
	if err := validator.ParseAndValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	addedUser, err := authenticationService.AddUser(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": addedUser,
		},
	})
}

func (auth *authentication) SignIn(c *fiber.Ctx) error {
	ctx := context.Background()

	request := &authenticationService.SignInDto{}
	if err := validator.ParseAndValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	tokens, err := authenticationService.VerifyUser(ctx, request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}
