package services

import (
	repository "SongLibrary/pkg/repositories"
	"SongLibrary/utils"
	"SongLibrary/utils/errs"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}
	if err := checkUserBlocked(user.ID); err != nil {
		return "", errs.ErrUserBlocked
	}
	accessToken, err = GenerateToken(user.ID, user.Username, user.RoleID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func checkUserBlocked(userID uint) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user.IsBlocked {
		return errs.ErrUserBlocked
	}
	return nil
}
