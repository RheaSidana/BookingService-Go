package seats

import "BookingService/src/model"

type ErrorResponse struct {
	Message string
}

type SeatBooking struct {
	SeatIdentifier string `json:"seat_identifiers" binding:"required"`
	IsBooked       bool   `json:"is_booked" binding:"required"`
}

type Seats struct {
	SeatClass string        `json:"seat_class" binding:"required"`
	Seat      []SeatBooking `json:"seats" binding:"required"`
}

type SeatsList struct {
	List []Seats `json:"seat_list" binding:"required"`
}

type Temp struct {
	Seat model.Seat
	SeatPrice model.SeatPricing
	TotalRows int `json:"total_rows" binding:"required"`
	TotalBooked int `json:"total_booked" binding:"required"`
}