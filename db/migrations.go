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
		&models.Song{},
	}

	for _, model := range migrateModels {
		err := dbConn.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}
	return nil
}
