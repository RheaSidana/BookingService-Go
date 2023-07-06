package model

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	BookingId string `gorm:"unique;not null" json:"booking_id" binding:"required"`
	UserId    uint `json:"user_id" binding:"required"`
	User      User   `gorm:"foreignKey:UserId; references:ID"`
	Total     int    `json:"total" binding:"required"`
}
