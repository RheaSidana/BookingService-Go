package booking

import "BookingService/src/model"

type ErrorResponse struct {
	Message string
}

type CreateBooking struct {
	Seats []model.SeatPrice `json:"seats" binding:"required"`
	Name  string            `json:"user_name" binding:"required"`
	Phone int               `json:"user_phone" binding:"required"`
}

type CreateBookingResponse struct {
	BookingId string  `json:"booking_id" binding:"required"`
	Total     float64 `json:"total_amount" binding:"required"`
}

type Booking struct {
	BookingId string            `json:"booking_id" binding:"required"`
	Total     float64           `json:"total_amount" binding:"required"`
	Seats     []model.SeatPrice `json:"seats" binding:"required"`
}

type Bookings struct{
	Bookings []Booking `json:"bookings" binding:"required"`
}