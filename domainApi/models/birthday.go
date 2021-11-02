package models

import "time"

type Birthday struct {
	ID        string `gorm:"primary_key"`
	UserID    string `json:"user_id" binding:"required"`
	ServerID  string `json:"server_id" binding:"required"`
	ChannelID string `json:"channel_id" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}
