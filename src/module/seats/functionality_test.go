package seats

import (
	"BookingService/src/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSeatBooking(t *testing.T) {
	mockSeat := model.Seat{
		SeatIdentifier: "A1",
	}

	mockSeatsBookedList := []string{"A1", "B1"}

	seatBooking := createSeatBooking(mockSeat, mockSeatsBookedList)

	assert.Equal(t, "A1", seatBooking.SeatIdentifier)
	assert.True(t, seatBooking.IsBooked)
}

func TestCreateSeats(t *testing.T) {
	mockSeat := model.Seat{
		SeatClass:     "Standard",
		SeatIdentifier: "A1",
	}

	mockSeatsBookedList := []string{"A1", "B1"}

	seats := createSeats(mockSeat, mockSeatsBookedList)

	assert.Equal(t, "Standard", seats.SeatClass)
	assert.Len(t, seats.Seat, 1)

	seatBooking := seats.Seat[0]
	assert.Equal(t, "A1", seatBooking.SeatIdentifier)
	assert.True(t, seatBooking.IsBooked)
}

func TestAppendSeats(t *testing.T) {
	mockSeat := model.Seat{
		SeatIdentifier: "A1",
		SeatClass:      "Standard",
	}

	mockSeatsBookedList := []string{"A1", "B1"}

	mockSeatWithClass := Seats{
		SeatClass: "Standard",
		Seat:      []SeatBooking{},
	}

	mockExpected := mockSeatWithClass
	mockExpected.Seat = []SeatBooking{
		{
			SeatIdentifier: "A1",
			IsBooked: true,
		},
	}

	seats := appendSeats(mockSeat, mockSeatsBookedList, mockSeatWithClass)

	assert.Len(t, seats.Seat, 1)
	assert.Equal(t,mockExpected, seats)
}

func TestCreateSeatsList(t *testing.T) {
	mockSeats := []model.Seat{
		{SeatIdentifier: "A1", SeatClass: "Standard"},
		{SeatIdentifier: "B1", SeatClass: "Standard"},
		{SeatIdentifier: "C1", SeatClass: "Premium"},
		{SeatIdentifier: "D1", SeatClass: "Premium"},
	}
	mockSeatsBooked := []model.SeatBooked{
		{SeatId: "A1", Price: 50.0},
		{SeatId: "B1", Price: 50.0},
		{SeatId: "D1", Price: 75.0},
	}

	seatsList := createSeatsList(mockSeats, mockSeatsBooked)

	assert.Len(t, seatsList.List, 2)

	standardSeats := seatsList.List[0]
	assert.Equal(t, "Standard", standardSeats.SeatClass)
	assert.Len(t, standardSeats.Seat, 2)
	assert.Equal(t, "A1", standardSeats.Seat[0].SeatIdentifier)
	assert.True(t, standardSeats.Seat[0].IsBooked)
	assert.Equal(t, "B1", standardSeats.Seat[1].SeatIdentifier)
	assert.True(t, standardSeats.Seat[1].IsBooked)

	premiumSeats := seatsList.List[1]
	assert.Equal(t, "Premium", premiumSeats.SeatClass)
	assert.Len(t, premiumSeats.Seat, 2)
	assert.Equal(t, "C1", premiumSeats.Seat[0].SeatIdentifier)
	assert.False(t, premiumSeats.Seat[0].IsBooked)
	assert.Equal(t, "D1", premiumSeats.Seat[1].SeatIdentifier)
	assert.True(t, premiumSeats.Seat[1].IsBooked)
}

func TestCreateSeatIdPrice(t *testing.T) {
	mockTemp := Temp{
		Seat: model.Seat{
			SeatIdentifier: "A1",
			SeatClass:      "Standard",
		},
		TotalBooked:  20,
		TotalRows:    100,
		SeatPrice: model.SeatPricing{
			MinPrice:     10.0,
			NormalPrice:  20.0,
			MaxPrice:     30.0,
		},
	}

	seatPrice := createSeatIdPrice(mockTemp)

	assert.Equal(t, "A1", seatPrice.Seat)
	assert.Equal(t, "Standard", seatPrice.Class)
	assert.Equal(t, 10.0, seatPrice.Price)

	//Update when min value not available or zero
	mockTemp.SeatPrice.MinPrice = 0
	seatPrice = createSeatIdPrice(mockTemp)

	assert.Equal(t, 20.0, seatPrice.Price)

	// Update the TotalBooked value to test the second condition
	mockTemp.TotalBooked = 50
	seatPrice = createSeatIdPrice(mockTemp)

	assert.Equal(t, "A1", seatPrice.Seat)
	assert.Equal(t, "Standard", seatPrice.Class)
	assert.Equal(t, 20.0, seatPrice.Price)

	//Update when normal value is not available or zero
	mockTemp.SeatPrice.NormalPrice = 0
	seatPrice = createSeatIdPrice(mockTemp)

	assert.Equal(t, 30.0, seatPrice.Price)

	// Update the TotalBooked value to test the third condition
	mockTemp.TotalBooked = 80
	seatPrice = createSeatIdPrice(mockTemp)

	// Assert the expected seat price fields based on the TotalBooked value
	assert.Equal(t, "A1", seatPrice.Seat)
	assert.Equal(t, "Standard", seatPrice.Class)
	assert.Equal(t, 30.0, seatPrice.Price)

	//Update when max value is not available or zero
	mockTemp.SeatPrice.NormalPrice = 50.0
	mockTemp.SeatPrice.MaxPrice = 0
	seatPrice = createSeatIdPrice(mockTemp)

	assert.Equal(t, 50.0, seatPrice.Price)

}
