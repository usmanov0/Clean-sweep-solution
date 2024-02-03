package domain

type UserRepository interface {
	Save(user *User) error
	UserExistByEmail(email string) (bool, error)
}
