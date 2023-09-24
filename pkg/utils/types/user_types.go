package types

type UserStatus uint64

const (
	Inactive UserStatus = iota
	Active
)

type UserRole uint64

const (
	Admin UserRole = iota + 1
	User
)
