package controllers

import (
	"context"
	"encoding/json"
	"errors"
	authenticationService "github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/services/authentication"
	validator "github.com/Shrijeeth/Golang-Fiber-App-Template/pkg"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type authentication struct{}

type AuthenticationControllerInterface interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	GoogleSignUp(c *fiber.Ctx) error
	GoogleCallback(c *fiber.Ctx) error
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

	userData := &types.AddUserData{
		Email:              request.Email,
		Password:           request.Password,
		UserRole:           request.UserRole,
		AuthenticationType: types.NormalAuthentication,
	}

	addedUser, err := authenticationService.AddUser(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": fiber.Map{
				"email": addedUser.Email,
			},
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

	userData := &types.VerifyUserData{
		Email:    request.Email,
		Password: request.Password,
	}
	tokens, err := authenticationService.VerifyUser(ctx, userData)
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

func (auth *authentication) GoogleSignUp(c *fiber.Ctx) error {
	googleConfig := configs.GetGoogleAuthConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	return c.Redirect(url, fiber.StatusFound)
}

func (auth *authentication) GoogleCallback(c *fiber.Ctx) error {
	ctx := context.Background()

	state := c.Query("state")
	if state != "randomstate" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "oauth states do not match",
		})
	}

	code := c.Query("code")
	googleConfig := configs.GetGoogleAuthConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "oauth code token exchange failed",
		})
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "oauth user data fetch failed",
		})
	}

	var userDetails types.GoogleOAuthData
	userDataResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "oauth user data parse failed",
		})
	}

	err = json.Unmarshal(userDataResponse, &userDetails)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "failed while parsing user data",
		})
	}

	err = validator.ValidateRequest(&userDetails)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "failed while parsing user data",
		})
	}

	existingUser, err := authenticationService.GetUserByEmail(userDetails.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || existingUser == nil {
			userData := &types.AddUserData{
				Email:              userDetails.Email,
				UserRole:           types.User,
				AuthenticationType: types.GoogleAuthentication,
			}
			addedUser, err := authenticationService.AddUser(userData)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			existingUser = addedUser
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
	}

	if existingUser.AuthenticationType == types.NormalAuthentication {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid Authentication Type",
		})
	}

	userData := &types.VerifyUserData{
		Email: existingUser.Email,
	}
	tokens, err := authenticationService.VerifyUser(ctx, userData)
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
