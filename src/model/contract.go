package model

type SeatPrice struct {
	Seat  string  `json:"seat" binding:"required"`
	Class  string  `json:"seat_class" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}