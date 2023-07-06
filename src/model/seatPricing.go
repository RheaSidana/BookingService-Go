package model

import (
	"gorm.io/gorm"
)

type SeatPricing struct {
	gorm.Model
	SeatClass   string  `gorm:"unique;not null" json:"seat_class" binding:"required"`
	MinPrice    float64 `json:"min_price" binding:"required"`
	MaxPrice    float64 `json:"max_price" binding:"required"`
	NormalPrice float64 `json:"normal_price" binding:"required"`
}
