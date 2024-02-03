package domain

import (
	"github.com/lib/pq"
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
	DeletedAt pq.NullTime
}
