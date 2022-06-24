package db_err

import (
	"errors"
	"gorm.io/gorm"
)

var (
	InternalError = errors.New("internal database error")

	RecordNotFound = gorm.ErrRecordNotFound
	WrongPassword  = errors.New("password hash mismatch")
)
