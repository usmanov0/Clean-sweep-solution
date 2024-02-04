package domain

type UserRepository interface {
	Save(user *User) error
	UserExistByEmail(email string) (bool, error)
	GetUsers() ([]User, error)
	FindById(userId int) (*User, error)
	GetHashedPasswordByEmail(email string) (string, error)
	UpdateUser(user *User) error
	DeleteUser(userId int) error
}
