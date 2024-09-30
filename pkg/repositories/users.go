package repository

import (
	"SongLibrary/db"
	"SongLibrary/logger"
	"SongLibrary/models"
	"SongLibrary/utils/errs"
	"errors"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = false").Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateError(err)
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ? AND is_deleted = false", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, TranslateError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("username = ? AND is_deleted = false", username).First(&user).Error

	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, TranslateError(err)
	}
	return &user, nil
}

func UserExists(username, email string) (bool, bool, error) {
	users, err := GetAllUsers()
	if err != nil {
		return false, false, err
	}

	var usernameExists, emailExists bool
	for _, user := range users {
		if user.Username == username {
			usernameExists = true
		}
		if user.Email == email {
			emailExists = true
		}
	}
	return usernameExists, emailExists, nil
}

func CreateUser(user models.User) (uint, error) {
	if err := db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return 0, TranslateError(err)
	}
	return user.ID, nil
}

func GetUserByUsernameAndPassword(username, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ? AND is_deleted = false", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, TranslateError(err)
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	updateData := map[string]interface{}{
		"full_name":  user.FullName,
		"username":   user.Username,
		"birth_date": user.BirthDate,
		"email":      user.Email,
		"password":   user.Password,
		"role_id":    user.RoleID,
	}
	for k, v := range updateData {
		if v == "" {
			delete(updateData, k)
		}
	}
	if len(updateData) == 0 {
		return nil
	}
	err := db.GetDBConn().Model(&user).Where("id = ? AND is_deleted = false", user.ID).Updates(updateData).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser] Failed to update user with ID %v: %v\n", user.ID, err)
		return TranslateError(err)
	}
	return nil
}

func DeleteUser(id uint) error {
	err := db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser] Failed to soft delete user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}

func UpdateUserPassword(id uint, newPassword string) error {
	err := db.GetDBConn().Model(&models.User{}).Where("id = ? AND is_deleted = false", id).Update("password", newPassword).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserPassword] Failed to update password for user with ID %v: %v\n", id, err)
		return TranslateError(err)
	}
	return nil
}
