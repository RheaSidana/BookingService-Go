package model

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	SeatIdentifier string `gorm:"unique;not null" json:"seat_identifier" binding:"required"`
	SeatClass string `gorm:"not null" json:"seat_class" binding:"required"`
	SeatPricing SeatPricing `gorm:"foreignKey:SeatClass; references:SeatClass"`
}
