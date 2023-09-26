package authentication

import "github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils/types"

type SignUpDto struct {
	Email    string         `json:"email" validate:"required,email,lte=255"`
	Password string         `json:"password" validate:"required,lte=255"`
	UserRole types.UserRole `json:"user_role" validate:"required"`
}

type SignInDto struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
