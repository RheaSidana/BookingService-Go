package operations

import (
	"BookingService/src/model"
	"fmt"
)

// func printSeat(seat model.Seat){
// 	fmt.Println("Seat:")
// 	fmt.Println("Seat Identifier : ", seat.SeatIdentifier)
// 	fmt.Println("Seat Class : ", seat.SeatClass)
// 	fmt.Println()
// }

func mapToSeat(index int, value string, seat model.Seat) model.Seat {
	switch index {
	case 1:
		seat.SeatIdentifier = value
	case 2:
		seat.SeatClass = value
	}

	return seat
}

func createSeat(record []string) model.Seat {

	var seat model.Seat

	for i, field := range record {
		seat = mapToSeat(i, string(field), seat)
	}

	return seat
}

func CreateSeat(records [][]string) []model.Seat {
	var seats []model.Seat

	for i, record := range records {
		if i == 0 {
			continue
		}
		seat := createSeat(record)
		// printSeat(seat)
		seats = append(seats, seat)
	}

	fmt.Println("............Create Seat Successfully............")

	return seats
}
