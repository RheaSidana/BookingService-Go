package seats

import (
	"BookingService/src/module/seatBooked"
	"BookingService/src/module/seatPricing"

	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(seatRepository Repository, 
	seatBookedRepo seatBooked.Repository, 
	seatPricingRepo seatPricing.Repository) Handler {
	return Handler{
		repository: seatRepository,
		seatBookedRepo: seatBookedRepo,
		seatPricingRepo: seatPricingRepo,
	}
}
