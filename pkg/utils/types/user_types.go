package types

type UserStatus uint64

const (
	InactiveUser UserStatus = iota
	ActiveUser
)

type UserRole uint64

const (
	Admin UserRole = iota + 1
	User
)

type GoogleOAuthData struct {
	Email string `json:"email" validate:"required,email,lte=255"`
	Name  string `json:"name" validate:"required,lte=255"`
}
