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
