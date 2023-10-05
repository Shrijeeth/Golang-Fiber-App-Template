package utils

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

type TokenMetadata struct {
	UserID            string
	Email             string
	UserRole          types.UserRole
	AdditionalDetails map[string]interface{}
	Expires           int64
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["id"].(string)

		// Expires time.
		expires := int64(claims["expires"].(float64))

		// User credentials.
		email := string(claims["email"].(string))
		userRole := claims["user_role"].(types.UserRole)
		additionalDetails := make(map[string]interface{})

		for key, value := range claims {
			additionalDetails[key] = value
		}

		return &TokenMetadata{
			UserID:            userID,
			Expires:           expires,
			Email:             email,
			UserRole:          userRole,
			AdditionalDetails: additionalDetails,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}