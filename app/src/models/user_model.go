package models

import (
	"github.com/Shrijeeth/Personal-Finance-Tracker-App/pkg/utils/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string           `gorm:"unique;not null"`
	PasswordHash string           `gorm:"not null"`
	UserStatus   types.UserStatus `gorm:"default:1"`
	UserRole     types.UserRole   `gorm:"default:2"`
}
