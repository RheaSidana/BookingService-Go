package model

type SeatPrice struct {
	Seat  string  `json:"seat" binding:"required"`
	Class  string  `json:"seat_class" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type CreateBooking struct {
	Seats []SeatPrice `json:"seats" binding:"required"`
	Name  string            `json:"user_name" binding:"required"`
	Phone int               `json:"user_phone" binding:"required"`
}

type CreateBookingResponse struct {
	BookingId string  `json:"booking_id" binding:"required"`
	Total     float64 `json:"total_amount" binding:"required"`
}

type BookingConfirm struct {
	BookingId string            `json:"booking_id" binding:"required"`
	Total     float64           `json:"total_amount" binding:"required"`
	Seats     []SeatPrice `json:"seats" binding:"required"`
}

type Bookings struct{
	Bookings []BookingConfirm `json:"bookings" binding:"required"`
}