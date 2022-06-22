package db_models

type User struct {
	ID   uint   `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
	Hash string `gorm:"column:hash"`
}
