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

type AuthenticationType uint64

const (
	NormalAuthentication AuthenticationType = iota + 1
	GoogleAuthentication
)

type GoogleOAuthData struct {
	Email string `json:"email" validate:"required,email,lte=255"`
	Name  string `json:"name" validate:"required,lte=255"`
}

type AddUserData struct {
	Username           string
	Email              string
	Password           string
	UserRole           UserRole
	AuthenticationType AuthenticationType
}

type VerifyUserData struct {
	Email    string `validate:"required,email,lte=255"`
	Password string
}
