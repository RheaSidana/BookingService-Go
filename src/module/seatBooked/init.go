package seatBooked

import (
	"gorm.io/gorm"
)


func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}