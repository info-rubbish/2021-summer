package database

import "errors"

var (
	ErrPasswordNotMatch = errors.New("password not match")
	ErrPermissionDenied = errors.New("permission not allowed")
)
