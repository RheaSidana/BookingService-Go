package booking

import (
	"BookingService/src/model"
	"BookingService/src/module/booking/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandlerWhenEmptyCreateBooking(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/booking", handler.Create)
	newCreateBooking := model.CreateBooking{}
	b, _ := json.Marshal(newCreateBooking)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/booking", body)
	respR := httptest.NewRecorder()
	expectedBooking := model.Booking{}
	repo.On("Create", newCreateBooking).Return(model.Booking{}, errors.New("error"))

	actualBooking, err := repo.Create(newCreateBooking)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, "error", err.Error())
}

func TestCreateHandlerWhenDbError(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/booking", handler.Create)
	newCreateBooking := model.CreateBooking{
		Seats: []model.SeatPrice{
			{
				Seat:  "A11",
				Class: "Standard",
				Price: 230.5,
			},
		},
		Name:  "abc",
		Phone: 9048334880,
	}
	b, _ := json.Marshal(newCreateBooking)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/booking", body)
	respR := httptest.NewRecorder()
	expectedBooking := model.Booking{}
	repo.On("Create", newCreateBooking).Return(model.Booking{}, errors.New("error"))

	actualBooking, err := repo.Create(newCreateBooking)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, "error", err.Error())
}

func TestCreateHandler(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/booking", handler.Create)
	newCreateBooking := model.CreateBooking{
		Seats: []model.SeatPrice{
			{
				Seat:  "A11",
				Class: "Standard",
				Price: 230.5,
			},
		},
		Name:  "abc",
		Phone: 9048334880,
	}
	b, _ := json.Marshal(newCreateBooking)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/booking", body)
	respR := httptest.NewRecorder()
	expectedBooking := model.Booking{
		BookingId: "BookingID",
		UserId:    23,
		Total:     231,
	}
	repo.On("Create", newCreateBooking).Return(expectedBooking, nil)

	actualBooking, err := repo.Create(newCreateBooking)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, nil, err)
}

func TestFindAllHandlerWhenEmptyQuery(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/bookings", handler.FindAll)
	// emailOrPhone := "anc@example.com"
	emailOrPhone := ""
	req, _ := http.NewRequest(http.MethodGet, "/bookings?userIdentifier="+emailOrPhone, nil)
	respR := httptest.NewRecorder()
	expectedBooking := model.Bookings{}
	repo.On("FindAll", emailOrPhone).Return(model.Bookings{}, errors.New("error"))

	actualBooking, err := repo.FindAll(emailOrPhone)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, "error", err.Error())
}

func TestFindAllHandlerWhenDBError(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/bookings", handler.FindAll)
	emailOrPhone := "anc@example.com"
	req, _ := http.NewRequest(http.MethodGet, "/bookings?userIdentifier="+emailOrPhone, nil)
	respR := httptest.NewRecorder()
	expectedBooking := model.Bookings{}
	repo.On("FindAll", emailOrPhone).Return(model.Bookings{}, errors.New("error"))

	actualBooking, err := repo.FindAll(emailOrPhone)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, "error", err.Error())
}

func TestFindAllHandler(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/bookings", handler.FindAll)
	emailOrPhone := "anc@example.com"
	req, _ := http.NewRequest(http.MethodGet, "/bookings?userIdentifier="+emailOrPhone, nil)
	respR := httptest.NewRecorder()
	expectedBooking := model.Bookings{
		Bookings: []model.BookingConfirm{
			{
				BookingId: "BookingID",
				Total: 200.0,
				Seats: []model.SeatPrice{
					{
						Seat: "SeatID",
						Class: "Class",
						Price: 200.0,
					},
				},
			},
		},
	}
	repo.On("FindAll", emailOrPhone).Return(expectedBooking, nil)

	actualBooking, err := repo.FindAll(emailOrPhone)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedBooking, actualBooking)
	assert.Equal(t, nil, err)
}