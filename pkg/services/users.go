package services

import (
	"SongLibrary/logger"
	"SongLibrary/models"
	repository "SongLibrary/pkg/repositories"
	"SongLibrary/utils"
	"SongLibrary/utils/errs"
)

func CreateUser(user models.User) (uint, error) {
	if err := user.ValidateCredentials(); err != nil {
		logger.Error.Printf("[service.CreateUser] validation error: %v\n", err)
		return 0, err
	}

	existingUser, err := repository.GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		logger.Error.Printf("[service.CreateUser] username already exists: %s\n", user.Username)
		return 0, errs.ErrUsernameExists
	}

	user.Password = utils.GenerateHash(user.Password)
	id, err := repository.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
