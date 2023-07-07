package booking 

import (
	"BookingService/src/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalSeatBookingTotal(t *testing.T) {
	mockSeatPrices := []model.SeatPrice{
		{Price: 10.0},
		{Price: 20.0},
		{Price: 15.0},
	}

	total := calSeatBookingTotal(mockSeatPrices)

	assert.Equal(t, 45.0, total)
}

func TestCreateBooking(t *testing.T) {
	count := int64(5)
	user := model.User{}
	user.ID = 10
	booking := model.CreateBooking{Seats: []model.SeatPrice{
		{
			Seat: "Seat",
			Class: "SeatClass",
			Price: 100.0,
		},
	}}
	total := 100.0

	expectedBookingId := "BK_6_U10_S1"
	expectedUserId := user.ID
	expectedTotal := 100

	bookSeats := createBooking(count, user, booking, total)

	assert.Equal(t, expectedBookingId, bookSeats.BookingId)
	assert.Equal(t, expectedUserId, bookSeats.UserId)
	assert.Equal(t, expectedTotal, bookSeats.Total)
}

func TestCreateSeatPrice(t *testing.T) {
	seat := model.SeatBooked{SeatId: "A1", Price: 50.0}
	seatClass := "Standard"

	expectedSeatPrice := model.SeatPrice{
		Seat: "A1", 
		Class: "Standard", 
		Price: 50.0,
	}

	seatPrice := createSeatPrice(seat, seatClass)

	assert.Equal(t, expectedSeatPrice, seatPrice)
}

func TestCreateBookingResp(t *testing.T) {
	bookingByUser := model.Booking{
		BookingId: "BK_1_U10_S3", 
		Total: 125,
	}
	seatsBooked := []model.SeatPrice{
		{Seat: "A1", Class: "Standard", Price: 50.0},
		{Seat: "B1", Class: "Premium", Price: 75.0},
	}

	expectedBooking := model.BookingConfirm{
		BookingId: "BK_1_U10_S3",
		Total:     125.0,
		Seats: []model.SeatPrice{
			{Seat: "A1", Class: "Standard", Price: 50.0},
			{Seat: "B1", Class: "Premium", Price: 75.0},
		},
	}

	booking := createBookingResp(bookingByUser, seatsBooked)

	assert.Equal(t, expectedBooking, booking)
}
