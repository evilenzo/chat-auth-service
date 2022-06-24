package db_models

type User struct {
	ID   uint64 `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
	Hash string `gorm:"column:hash"`
}

type Client struct {
	ID    uint64 `gorm:"column:user_id"`
	Token string `gorm:"column:token"`
}
