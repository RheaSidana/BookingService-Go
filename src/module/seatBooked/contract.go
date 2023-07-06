package seatBooked

type ErrorResponse struct {
	Message string
}

type SeatBooking struct{
	SeatIdentifier string `json:"seat_identifiers" binding:"required"`
	IsBooked bool `json:"is_booked" binding:"required"`
}

type Seats struct{
	SeatClass string  `json:"seat_class" binding:"required"`
	SeatIdentifiers []string  `json:"seat_identifiers" binding:"required"`
}

type SeatsList struct{
	List []Seats `json:"seat_list" binding:"required"`
}