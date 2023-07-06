package seatBooked

import "BookingService/src/model"

func CreateSeatsBookedList(seats []model.SeatBooked) ([]string) {
	var seatsList []string

	for	_, seat := range seats{
		seatsList = append(seatsList, seat.SeatId)
	}

	return seatsList
}