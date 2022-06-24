package db_operator

import (
	"errors"
	model "main/db_operator/db_models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Internal database error: ", err)
		return exists, err
	}

	return exists, nil
}
