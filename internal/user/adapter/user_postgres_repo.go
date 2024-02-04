package adapter

import (
	"Clean-Sweep-Solutions_/internal/user/domain"
	"Clean-Sweep-Solutions_/pkg/errors"
	"github.com/jackc/pgx"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) domain.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) Save(user *domain.User) error {
	queryStatement :=
		`INSERT INTO users(full_name, email, phone, password, role, created_at, updated_at) 
		VALUES ($1,$2,$3,$4)`

	_, err := u.db.Exec(queryStatement, user.FullName, user.Email, user.Phone,
		user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return errors.ErrFailedExecQuery
	}

	return nil
}

func (u *userRepo) UserExistByEmail(email string) (bool, error) {
	var exist bool
	queryStatement :=
		`SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE email = $1
		)`

	err := u.db.QueryRow(queryStatement, email).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (u *userRepo) GetUsers() ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) FindById(userId int) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) UpdateUser(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) DeleteUser(userId int) error {
	//TODO implement me
	panic("implement me")
}
