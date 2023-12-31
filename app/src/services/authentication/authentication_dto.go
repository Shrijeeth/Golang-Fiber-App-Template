package authentication

import "github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"

type SignUpDto struct {
	Username string         `json:"username" validate:"required,lte=255"`
	Email    string         `json:"email" validate:"required,email,lte=255"`
	Password string         `json:"password" validate:"required,lte=255"`
	UserRole types.UserRole `json:"user_role" validate:"required"`
}

type SignInDto struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
