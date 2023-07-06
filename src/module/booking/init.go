package booking

import (
	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(seatRepository Repository) Handler {
	return Handler{repository: seatRepository}
}
