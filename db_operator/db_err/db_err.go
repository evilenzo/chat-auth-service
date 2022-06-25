package db_err

import (
	"errors"
)

var (
	InternalError = errors.New("internal database error")

	RecordNotFound = errors.New("record not found")
	WrongPassword  = errors.New("password hash mismatch")
)
