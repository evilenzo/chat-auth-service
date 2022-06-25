package db_operator

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	dbe "main/db_operator/db_err"
	model "main/db_operator/db_models"
	"main/db_operator/encryption"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type DatabaseOperator struct {
	db *bun.DB
}

func CreateDatabaseOperator(db *bun.DB) DatabaseOperator {
	return DatabaseOperator{db}
}

func idToCodeword(id int64) []byte {
	bs := []byte(strconv.FormatInt(id, 10))
	return bs
}

// API

func (dbo *DatabaseOperator) NameExists(ctx context.Context, name string) (bool, error) {
	exists, err := dbo.db.NewSelect().
		Model((*model.User)(nil)).
		Where("name = ?", name).
		Exists(ctx)

	if err != nil {
		log.Error("Internal database error in dbo::NameExists: ", err)
		return false, dbe.InternalError
	}

	return exists, nil
}

func (dbo *DatabaseOperator) AuthApp(ctx context.Context, id int64, password string) (string, error) {
	var user model.User

	err := dbo.db.NewSelect().
		Model(&user).
		Where("user_id = ?", id).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Trace("Tried to query not existing user with id ", id)
			return "", dbe.RecordNotFound
		}

		log.Error("Internal database error in dbo::AuthApp: ", err)
		return "", dbe.InternalError
	}

	if encryption.VerifyPassword(password, user.Hash) {
		codeword := idToCodeword(id)
		client := model.Client{ID: id, Token: encryption.GetToken(codeword)}
		_, err := dbo.db.NewInsert().
			Model(&client).
			Exec(ctx)

		if err != nil {
			log.Error("Error during token insertion in dbo::AuthApp: ", err)
			return "", dbe.InternalError
		}

		return client.Token, nil
	}

	return "", dbe.WrongPassword
}
