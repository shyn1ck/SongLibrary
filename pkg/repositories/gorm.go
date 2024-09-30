package repository

import (
	"SongLibrary/utils/errs"
	"errors"
	"gorm.io/gorm"
)

func TranslateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	return err
}
