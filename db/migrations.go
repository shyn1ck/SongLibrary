package db

import (
	"SongLibrary/models"
	"errors"
	"fmt"
)

func Migrate() (err error) {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	migrateModels := []interface{}{
		&models.User{},
		&models.Song{},
		&models.Role{},
		&models.Album{},
		&models.Artist{},
	}

	for _, model := range migrateModels {
		err := dbConn.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}

	err = SeedRoles()
	if err != nil {
		return fmt.Errorf("failed to seed roles: %v", err)
	}

	return nil
}

func SeedRoles() error {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "listener"},
		{Name: "artist"},
	}

	for _, role := range roles {
		if err := GetDBConn().FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
