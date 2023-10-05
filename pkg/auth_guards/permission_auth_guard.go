package auth_guards

import (
	"os"
	"strconv"
	"time"

	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/middleware"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func PermissionAuthGuard(allowedPermissions []types.UserRole) func(*fiber.Ctx) error {
	jwtConfig := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: middleware.JwtError,
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token, err := utils.ExtractTokenMetadata(ctx)
			now := time.Now().Unix()
			if now > token.Expires {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"error":   "Access Token Expired",
				})
			}
			if err == nil {
				userID, err := strconv.Atoi(token.UserID)
				if err != nil {
					return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"success": false,
						"error":   err.Error(),
					})
				}

				if len(allowedPermissions) == utils.IntZero {
					ctx.Locals("userID", uint(userID))
					return ctx.Next()
				}

				for _, permission := range allowedPermissions {
					if token.UserRole == permission {
						ctx.Locals("userID", uint(userID))
						return ctx.Next()
					}
				}
			}
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized",
			})
		},
	}

	return middleware.JWTProtected(jwtConfig)
}
