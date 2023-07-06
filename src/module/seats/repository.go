package seats

import (
	"BookingService/src/model"
	"strconv"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Seat, error)
	FindSeat(id string) (model.Seat, error)
	FindSeatClass(id string) (string, error)
	Count(sclass string) (int, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) FindSeatClass(id string) (string, error) {
	var seatClass string
	if res := r.client.Model(&model.Seat{}).Select(
		"seat_class").Where(
		"seat_identifier=?", id).Find(
		&seatClass); res.Error != nil {
		return "", res.Error
	}
	return seatClass, nil
}

func (r *repository) FindAll() ([]model.Seat, error) {
	var seats []model.Seat
	if res := r.client.Order("seat_class asc").Find(&seats); res.Error != nil {
		return []model.Seat{}, res.Error
	}

	return seats, nil
}

func (r *repository) FindSeat(id string) (model.Seat, error) {
	ID, _ := strconv.Atoi(id)

	var seatPricing model.Seat
	if res := r.client.Where("id=?", ID).Find(&seatPricing); res.Error != nil {
		return model.Seat{}, res.Error
	}

	return seatPricing, nil
}

func (r *repository) Count(sclass string) (int, error) {
	var seats []model.Seat

	res := r.client.Where("seat_class=?", sclass).Find(&seats)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
