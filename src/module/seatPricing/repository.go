package seatPricing

import (
	"BookingService/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	Find(id string) (model.SeatPricing, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Find(id string) (model.SeatPricing, error){ 
	var seatPricing model.SeatPricing
	if res := r.client.Where("seat_class=?",id).Find(&seatPricing); res.Error != nil {
		return model.SeatPricing{}, res.Error
	}

	return seatPricing, nil
}