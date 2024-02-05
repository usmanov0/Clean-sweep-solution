package domain

import "example.com/m/internal/genproto/user_pb/pb"

type UserRepository interface {
	Save(user *User) error
	UserExistByEmail(email string) (bool, error)
	GetUsers() ([]pb.User, error)
	FindById(id *pb.UserId) (*pb.User, error)
	GetHashedPasswordByEmail(email string) (string, error)
	UpdateUser(user *pb.UserUpdate) error
	DeleteUser(id *pb.UserId) error
}
