package app

import (
	"example.com/m/internal/user/domain"
	"example.com/m/pkg/errors"
	"example.com/m/pkg/utils"
	"fmt"
	"log"
	"time"
)

type userUseCase struct {
	userRepo    domain.UserRepository
	userFactory domain.UserFactory
}

type UserUseCase interface {
	SignUpAdmin(user *domain.NewUser) error
	SignUpUser(user *domain.NewUser) error
	SignInUser(email, password string) (bool, error)
	GetUsers() ([]domain.User, error)
	GetUser(userId int) (*domain.User, error)
	Update(user *domain.User) error
	Delete(userId int) error
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
		fmt.Sprintf("failed sign up admin")
		return err
	}

	exists, err := u.userRepo.UserExistByEmail(adminUser.Email)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return err
	}

	if exists {
		return errors.ErrEmailExist
	} else {
		err := u.userRepo.Save(adminUser)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *userUseCase) SignUpUser(user *domain.NewUser) error {
	userSignUp := u.userFactory.CreateUser(user)

	err := utils.ValidateUserInfoForSignIn(
		userSignUp.Email,
		userSignUp.Password,
	)

	if err != nil {
		fmt.Sprintf("failed sign up user")
		return err
	}

	exists, err := u.userRepo.UserExistByEmail(userSignUp.Email)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return err
	}

	if exists {
		return errors.ErrEmailExist
	} else {
		err := u.userRepo.Save(userSignUp)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *userUseCase) SignInUser(email, password string) (bool, error) {
	err := utils.ValidateUserInfoForSignIn(email, password)

	if err != nil {
		return false, nil
	}

	ok, err := u.userRepo.UserExistByEmail(email)
	if !ok || err != nil {
		return false, errors.ErrUserNotFound
	}

	return true, nil
}

func (u *userUseCase) GetUsers() ([]domain.User, error) {
	userList, err := u.userRepo.GetUsers()

	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return userList, nil
}

func (u *userUseCase) GetUser(userId int) (*domain.User, error) {
	user, err := u.userRepo.FindById(userId)

	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCase) Update(userUpdate *domain.User) error {
	var id int
	existUser, err := u.userRepo.FindById(id)

	if err != nil {
		return errors.ErrUserNotFound
	}
	existUser.FullName = userUpdate.FullName
	existUser.Phone = userUpdate.Phone
	existUser.Password = userUpdate.Password
	existUser.UpdatedAt = time.Now()

	err = u.userRepo.UpdateUser(existUser)
	if err != nil {
		return errors.ErrUpdateFailed
	}

	return nil
}

func (u *userUseCase) Delete(userId int) error {
	err := u.userRepo.DeleteUser(userId)

	if err != nil {
		return errors.ErrUserDeleteFailed
	}

	return nil
}
