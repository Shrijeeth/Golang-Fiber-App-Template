package models

import (
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username           string                   `gorm:"null"`
	Email              string                   `gorm:"unique;not null"`
	PasswordHash       string                   `gorm:"null"`
	UserStatus         types.UserStatus         `gorm:"default:1"`
	UserRole           types.UserRole           `gorm:"default:2"`
	AuthenticationType types.AuthenticationType `gorm:"not null;default:1"`
}
