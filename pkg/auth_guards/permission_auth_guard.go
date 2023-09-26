package auth_guards

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/middleware"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
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
				for _, permission := range allowedPermissions {
					if token.UserRole == permission {
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
