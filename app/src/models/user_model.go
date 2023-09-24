package models

import (
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	PasswordHash string
	UserStatus   types.UserStatus
	UserRole     types.UserRole
}
