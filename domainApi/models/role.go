package models

type Role struct {
	Id   uint64 `gorm:"primary_key"`
	Name string `gorm:"uniqueIndex" json:"name"`
}
