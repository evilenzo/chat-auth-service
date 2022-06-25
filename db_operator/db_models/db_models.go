package db_models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID   int64  `bun:"user_id,pk,autoincrement"`
	Name string `bun:"name"`
	Hash string `bun:"hash"`
}

type Client struct {
	bun.BaseModel `bun:"table:clients"`

	ID    int64  `bun:"user_id"`
	Token string `bun:"token"`
}
