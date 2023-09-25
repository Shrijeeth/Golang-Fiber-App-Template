package auth_guards

import (
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/middleware"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils"
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils/types"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func PermissionAuthGuard(allowedPermissions []types.UserRole) func(*fiber.Ctx) error {
	jwtConfig := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: middleware.JwtError,
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token := ctx.Locals("user").(*jwt.Token)
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if role, found := claims["user_role"].(types.UserRole); found {
					if len(allowedPermissions) == utils.IntZero {
						return ctx.Next()
					}
					for _, permission := range allowedPermissions {
						if role == permission {
							return ctx.Next()
						}
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
