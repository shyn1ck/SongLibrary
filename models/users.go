package models

import (
	"SongLibrary/utils/errs"
	"strings"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
	Username  string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	RoleID    uint      `json:"role_id" gorm:"not null"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
	IsBlocked bool      `json:"-" gorm:"type:bool;not null;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(100);unique;not null"`
}

func (u User) ValidateCredentials() (err error) {
	if strings.TrimSpace(u.Username) == "" {
		return errs.ErrValidationFailed
	}

	if strings.TrimSpace(u.Email) == "" {
		return errs.ErrValidationFailed
	}

	if u.RoleID == 0 {
		return errs.ErrValidationFailed
	}

	if len(u.Password) < 8 {
		return errs.ErrValidationFailed
	}

	switch u.RoleID {
	case 1:
		return errs.ErrPermissionDenied
	case 2, 3:
	default:
		return errs.ErrValidationFailed
	}
	return nil
}
