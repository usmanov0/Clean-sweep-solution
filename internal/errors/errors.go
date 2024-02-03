package errors

// User errors
type UserError string

func (u UserError) Error() string {
	panic("implement me")
}

const (
	ErrFailedExecQuery UserError = "failed to execute query"
	ErrUserNotFound    UserError = "user not found"
)
