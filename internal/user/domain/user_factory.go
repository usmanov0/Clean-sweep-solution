package domain

import (
	"time"
)

type UserFactory struct{}

func (f UserFactory) CreateAdmin(user *NewUser) *User {
	return &User{
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		Role:      "admin",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
}

func (f UserFactory) CreateUser(user *NewUser) *User {
	return &User{
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		Role:      "user",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
}
