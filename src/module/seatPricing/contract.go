package seatPricing

type Seat struct {
	Seat  string  `json:"seat" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
