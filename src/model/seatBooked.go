package model

import (
	"gorm.io/gorm"
)

type SeatBooked struct {
	gorm.Model
	BookingId string  `json:"booking_id" binding:"required"`
	Booking   Booking `gorm:"foreignKey:BookingId; reference:BookingId"`
	SeatId    string  `json:"seat_id" binding:"required"`
	Seat      Seat    `gorm:"foreignKey:SeatId; references:SeatIdentifier"`
	Price     float64 `json:"price" binding:"required"`
}
