package seats

import (
	"BookingService/src/model"
	seatsBookedMock "BookingService/src/module/seatBooked/mocks"
	seatPricingMock "BookingService/src/module/seatPricing/mocks"
	"BookingService/src/module/seats/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindAllHandlerWhenDbError(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/seats", handler.FindAllSeats)
	req, _ := http.NewRequest(http.MethodGet, "/seats", nil)
	respR := httptest.NewRecorder()
	expectedSeats := []model.Seat{}
	repo.On("FindAll").Return(expectedSeats, errors.New("error"))

	actualSeats, err := repo.FindAll()
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedSeats, actualSeats)
	assert.Equal(t, "error", err.Error())
}

func TestFindAllSeatsHandlerWhenDbErrorSeatsBooked(t *testing.T) {
	repo := new(mocks.Repository)
	repoSeatsBooked := new(seatsBookedMock.Repository)
	handler := Handler{
		repository: repo, 
		seatBookedRepo: repoSeatsBooked,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/seats", handler.FindAllSeats)
	req, _ := http.NewRequest(http.MethodGet, "/seats", nil)
	respR := httptest.NewRecorder()
	expectedSeats := []model.Seat{
		{
			SeatIdentifier: "SeatID",
		},
	}
	repo.On("FindAll").Return(expectedSeats, nil)

	expectedSeatsBooked := []model.SeatBooked{}
	repoSeatsBooked.On("FindAll").Return(expectedSeatsBooked, errors.New("error"))

	actualSeats, err := repo.FindAll()
	actualSeatsBooked, er := repoSeatsBooked.FindAll()
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedSeats, actualSeats)
	assert.Equal(t, expectedSeatsBooked, actualSeatsBooked)
	assert.Equal(t, nil, err)
	assert.Equal(t, "error", er.Error())
}

func TestFindAllSeatsHandler(t *testing.T) {
	repo := new(mocks.Repository)
	repoSeatsBooked := new(seatsBookedMock.Repository)
	handler := Handler{
		repository: repo, 
		seatBookedRepo: repoSeatsBooked,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/seats", handler.FindAllSeats)
	req, _ := http.NewRequest(http.MethodGet, "/seats", nil)
	respR := httptest.NewRecorder()
	expectedSeats := []model.Seat{
		{
			SeatIdentifier: "SeatID",
		},
	}
	repo.On("FindAll").Return(expectedSeats, nil)

	expectedSeatsBooked := []model.SeatBooked{
		{
			BookingId: "BookingID",
		},
	}
	repoSeatsBooked.On("FindAll").Return(expectedSeatsBooked, nil)

	actualSeats, err := repo.FindAll()
	actualSeatsBooked, er := repoSeatsBooked.FindAll()
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedSeats, actualSeats)
	assert.Equal(t, expectedSeatsBooked, actualSeatsBooked)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, er)
}

func TestFindSeatHandler(t *testing.T) {
	repo := new(mocks.Repository)
	repoSeatsBooked := new(seatsBookedMock.Repository)
	repoSeatPricing := new(seatPricingMock.Repository)
	handler := Handler{
		repository: repo, 
		seatPricingRepo: repoSeatPricing,
		seatBookedRepo: repoSeatsBooked,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/seats/:id", handler.FindSeat)
	req, _ := http.NewRequest(http.MethodGet, "/seats/1", nil)
	respR := httptest.NewRecorder()
	expectedSeats := model.Seat{
		SeatIdentifier: "S1",
		SeatClass: "C1",
	}
	expectedSeatPrice := model.SeatPricing{
		SeatClass: "C1",
		MinPrice: 230.2,
		NormalPrice: 250.6,
		MaxPrice: 300.3,
	}
	rowCount := 100
	bookedCount := 50
	repo.On("FindSeat", "1").Return(expectedSeats, nil)
	repoSeatPricing.On("Find", expectedSeats.SeatClass).Return(expectedSeatPrice, nil)
	repo.On("Count", expectedSeats.SeatClass).Return(rowCount, nil)
	repoSeatsBooked.On("Count", expectedSeats.SeatClass).Return(bookedCount, nil)

	actualSeats, err := repo.FindSeat("1")
	actualSeatPrice, er := repoSeatPricing.Find(expectedSeatPrice.SeatClass)
	actualRcount, e := repo.Count(expectedSeatPrice.SeatClass)
	actualBCount, erro := repoSeatsBooked.Count(expectedSeatPrice.SeatClass)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedSeats, actualSeats)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedSeatPrice, actualSeatPrice)
	assert.Equal(t, nil, er)
	assert.Equal(t, rowCount, actualRcount)
	assert.Equal(t, nil, e)
	assert.Equal(t, bookedCount, actualBCount)
	assert.Equal(t, nil, erro)
}