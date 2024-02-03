package app

import "Clean-Sweep-Solutions_/internal/user/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

type UserUseCase interface {
	SignUpAdmin(user *domain.User) error
	SignUpUser(user *domain.User) error
}

func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return userUseCase{userRepo: userRepo}
}

func (u userUseCase) SignUpAdmin(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) SignUpUser(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
