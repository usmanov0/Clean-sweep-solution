package adapter

import (
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/domain"
	"example.com/m/pkg/errors"
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

func (u *userRepo) GetUsers() ([]pb.User, error) {
	queryStatement := `
		SELECT u.id, u.full_name, u.email, u.phone, u.role
		FROM users u 
	`

	rows, err := u.db.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	var users []pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Phone, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepo) FindById(userId *pb.UserId) (*pb.User, error) {
	queryStatement := `
		SELECT u.id, u.full_name, u.email, u.phone, u.password, u.role, u.created_at, u.updated_dt, u.deleted_at
		FROM users u
		WHERE u.id = $1
	`
	var user pb.User
	err := u.db.QueryRow(queryStatement, userId).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetHashedPasswordByEmail(email string) (string, error) {
	var hashedPassword string
	queryStatement := `
		SELECT password
		FROM users
		WHERE email = $1`

	err := u.db.QueryRow(queryStatement, email).Scan(&hashedPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", errors.ErrUserNotFound
		}
		return "", err
	}
	return hashedPassword, nil
}

func (u *userRepo) UpdateUser(user *pb.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) DeleteUser(id *pb.UserId) error {
	//TODO implement me
	panic("implement me")
}
