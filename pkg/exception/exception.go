package exception

import "errors"

var (
	EmailAlreadyExists   = errors.New("email already exists")
	EmailOrPasswordWrong = errors.New("email or password wrong")
	PasswordWrong        = errors.New("password wrong")

	UserNotFound = errors.New("user not found")
)
