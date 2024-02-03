package domain

import (
	"time"
)

type UserRole string

const (
	adminRole UserRole = "admin"
	userRole  UserRole = "user"
)

type User struct {
	Id        int
	FullName  string
	Email     string
	Phone     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type NewUser struct {
	FullName string
	Email    string
	Phone    string
	Password string
}
