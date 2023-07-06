package booking

import (
	"BookingService/src/model"
	"math"
	"strconv"
)

func calSeatBookingTotal(seats []model.SeatPrice) float64 {
	sum := 0.0
	for _, seat := range seats {
		sum += seat.Price
	}

	return sum
}

func createBooking(count int64, user model.User, booking CreateBooking, total float64) model.Booking {
	rows := strconv.Itoa(int(count + 1))
	uid := strconv.Itoa(int(user.ID))
	length := strconv.Itoa(len(booking.Seats))

	//create booking
	bookSeats := model.Booking{
		BookingId: "BK_" + rows + "_U" + uid + "_S" + length,
		UserId:    user.ID,
		Total:     int(math.Ceil(total)),
	}

	return bookSeats
}

func createSeatPrice(seat model.SeatBooked, seatClass string) model.SeatPrice {
	return model.SeatPrice{
		Seat:  seat.SeatId,
		Class: seatClass,
		Price: seat.Price,
	}
}

func createBookingResp(bookingByUser model.Booking, seatsBooked []model.SeatPrice) Booking {
	var booking Booking
	booking.BookingId = bookingByUser.BookingId
	booking.Total = float64(bookingByUser.Total)
	booking.Seats = seatsBooked
	return booking
}
