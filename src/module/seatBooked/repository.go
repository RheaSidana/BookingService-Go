package seatBooked

import (
	"BookingService/src/model"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.SeatBooked, error)
	FindAllByBookingId(booking model.Booking) ([]model.SeatBooked, error)
	Count(sclass string) (int, error)
	CheckBookingSeat(seats []model.SeatPrice) (bool, error)
	Create(booking model.Booking, seats []model.SeatPrice) (bool, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) FindAllByBookingId(booking model.Booking) ([]model.SeatBooked, error) {
	var seats []model.SeatBooked
	if res := r.client.Where(
		"booking_id=?", booking.BookingId).Find(
		&seats); res.Error != nil {
		return []model.SeatBooked{}, res.Error
	}
	return seats, nil
}

func (r *repository) Create(booking model.Booking, seats []model.SeatPrice) (bool, error) {
	for _, seat := range seats {
		seatsBooked := model.SeatBooked{
			BookingId: booking.BookingId,
			SeatId:    seat.Seat,
			Price:     seat.Price,
		}
		//add to seatbooked
		if res := r.client.Create(&seatsBooked); res.Error != nil {
			return false, res.Error
		}
	}
	return true, nil
}

func (r *repository) FindAll() ([]model.SeatBooked, error) {
	var seats []model.SeatBooked
	if res := r.client.Find(&seats); res.Error != nil {
		return []model.SeatBooked{}, res.Error
	}

	return seats, nil
}

func (r *repository) Count(sclass string) (int, error) {
	var count int

	res := r.client.Model(&model.Seat{}).Select("count(*)").Joins("left join seat_bookeds sb on seats.seat_identifier = sb.seat_id").Where("seats.seat_identifier in (Select seat_id from seat_bookeds)").Scan(&count)
	if res.Error != nil {
		return 0, res.Error
	}

	return count, nil
}

func (r *repository) CheckBookingSeat(seats []model.SeatPrice) (bool, error) {
	booked := ""
	counting := 0
	for _, seat := range seats {
		var count int64 = 0
		if res := r.client.Where(
			"seat_id=?", seat.Seat,
		).Find(&model.SeatBooked{}).Count(&count); res.Error != nil {
			return true, res.Error
		}
		if count > 0 {
			counting++
			booked += " seat : " + seat.Seat + "\n"
		}
	}
	if counting > 0 {
		return true, errors.New(booked)
	}
	return false, nil
}
