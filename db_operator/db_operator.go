package db_operator

import (
	"errors"
	"gorm.io/gorm"
	"log"
	model "main/db_operator/db_models"
)

type DatabaseOperator struct {
	db *gorm.DB
}

func CreateDatabaseOperator(db *gorm.DB) DatabaseOperator {
	return DatabaseOperator{db}
}

func (dbo *DatabaseOperator) NameExists(name string) (bool, error) {
	var exists bool
	err := dbo.db.Model(model.User{}).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error

	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Internal database error: %v", err)
		return exists, err
	}

	return exists, nil
}
