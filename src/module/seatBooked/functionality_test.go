package seatBooked

import (
	"BookingService/src/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSeatsBookedList(t *testing.T) {
	mockSeatBookings := []model.SeatBooked{
		{SeatId: "A1"},
		{SeatId: "B1"},
		{SeatId: "C1"},
	}

	seatsList := CreateSeatsBookedList(mockSeatBookings)

	expectedSeatsList := []string{"A1", "B1", "C1"}

	assert.Equal(t, expectedSeatsList, seatsList)
}