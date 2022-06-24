package db_operator

import (
	"encoding/binary"
	"errors"
	dbe "main/db_operator/db_err"
	model "main/db_operator/db_models"
	"main/db_operator/encryption"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DatabaseOperator struct {
	db *gorm.DB
}

func CreateDatabaseOperator(db *gorm.DB) DatabaseOperator {
	return DatabaseOperator{db}
}

func idToCodeword(id uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, id)
	return bs
}

// API

func (dbo *DatabaseOperator) NameExists(name string) (bool, error) {
	var exists bool
	err := dbo.db.Model(model.User{}).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Internal database error in dbo::NameExists: ", err)
		return false, dbe.InternalError
	}

	return exists, nil
}

func (dbo *DatabaseOperator) AuthApp(id uint64, password string) (string, error) {
	var user model.User

	// BUG: Wrong password if not found
	if err := dbo.db.Limit(1).Find(&user, id).Error; err == nil {
		if encryption.VerifyPassword(password, user.Hash) {
			codeword := idToCodeword(id)
			client := model.Client{id, encryption.GetToken(codeword)}
			err := dbo.db.Create(&client).Error
			if err == nil {
				return client.Token, nil
			}
			log.Error("Error during token insertion in dbo::AuthApp: ", err)
			return "", dbe.InternalError
		}
		return "", dbe.WrongPassword
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", dbe.RecordNotFound
	} else {
		log.Error("Internal database error in dbo::AuthApp: ", err)
		return "", dbe.InternalError
	}
}
