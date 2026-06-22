package domain

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")
)

var (
	ErrJobNotFound = errors.New("Job not found")
)
