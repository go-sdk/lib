package db

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

func IsErrNoRows(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows)
}
