package app

import (
	"Clean-Sweep-Solutions_/internal/user/domain"
	"Clean-Sweep-Solutions_/internal/utils"
	"errors"
	"fmt"
	"log"
)

type userUseCase struct {
	userRepo    domain.UserRepository
	userFactory domain.UserFactory
}

type UserUseCase interface {
	SignUpAdmin(user *domain.NewUser) error
	SignUpUser(user *domain.NewUser) error
}

func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) SignUpAdmin(admin *domain.NewUser) error {
	adminUser := u.userFactory.CreateAdmin(admin)

	err := utils.ValidateUserInfoForSignUp(
		adminUser.FullName,
		adminUser.Email,
		adminUser.Phone,
		adminUser.Password,
	)

	if err != nil {
		fmt.Sprintf("can't sign up admin")
	}

	exists, err := u.userRepo.UserExistByEmail(adminUser.Email)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return err
	}

	if exists {
		return errors.New("email already exists")
	} else {
		err := u.userRepo.Save(adminUser)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *userUseCase) SignUpUser(user *domain.NewUser) error {
	return nil
}
