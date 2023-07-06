package seats

import (
	"BookingService/src/model"
	"BookingService/src/module/seatBooked"

	"golang.org/x/exp/slices"
)

func createSeatBooking(seat model.Seat, seatsBookedList []string) SeatBooking {
	var seatBooking SeatBooking
	seatBooking.SeatIdentifier = seat.SeatIdentifier
	if slices.Contains(seatsBookedList, seat.SeatIdentifier) {
		seatBooking.IsBooked = true
	}
	return seatBooking
}

func createSeats(seat model.Seat, seatsBookedList []string) Seats {
	var seatWithClass Seats
	seatWithClass.SeatClass = seat.SeatClass
	seatBooking := createSeatBooking(seat, seatsBookedList)
	seatWithClass.Seat = append(seatWithClass.Seat, seatBooking)
	return seatWithClass
}

func appendSeats(seat model.Seat, seatsBookedList []string, seatWithClass Seats) Seats {
	seatBooking := createSeatBooking(seat, seatsBookedList)
	seatWithClass.Seat = append(seatWithClass.Seat, seatBooking)
	return seatWithClass
}

func createSeatsList(seats []model.Seat, seatsBooked []model.SeatBooked) SeatsList {
	var seatsList SeatsList

	seatsBookedList := seatBooked.CreateSeatsBookedList(seatsBooked)

	var seatWithClass Seats
	flag := 0
	sclass := ""
	for _, seat := range seats {
		if sclass != seat.SeatClass {
			if flag >= 1 {
				seatsList.List = append(seatsList.List, seatWithClass)
			}
			seatWithClass = createSeats(seat, seatsBookedList)
			flag++
			sclass = seat.SeatClass
		} else {
			seatWithClass = appendSeats(seat, seatsBookedList, seatWithClass)
		}
	}
	seatsList.List = append(seatsList.List, seatWithClass)

	return seatsList
}

func createSeatIdPrice(temp Temp) model.SeatPrice {
	seat := model.SeatPrice{
		Seat:  temp.Seat.SeatIdentifier,
		Class: temp.Seat.SeatClass,
	}

	if temp.TotalBooked < int(temp.TotalRows*40/100) {
		if temp.SeatPrice.MinPrice != 0 {
			seat.Price = temp.SeatPrice.MinPrice
		} else {
			seat.Price = temp.SeatPrice.NormalPrice
		}
	}else if temp.TotalBooked >= int(temp.TotalRows*40/100) && temp.TotalBooked <= int(temp.TotalRows*60/100){
		if temp.SeatPrice.NormalPrice != 0 {
			seat.Price = temp.SeatPrice.NormalPrice
		} else {
			seat.Price = temp.SeatPrice.MaxPrice
		}
	}else{
		if temp.SeatPrice.MaxPrice != 0 {
			seat.Price = temp.SeatPrice.MaxPrice
		} else {
			seat.Price = temp.SeatPrice.NormalPrice
		}
	}

	return seat
}
