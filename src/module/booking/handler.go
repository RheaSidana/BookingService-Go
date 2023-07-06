package booking

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
}

func (h *Handler) Create(c *gin.Context){
	var booking CreateBooking
	c.BindJSON(&booking)

	if reflect.DeepEqual(booking, CreateBooking{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to create booking."})
		return
	}

	bookingConfirm, err := h.repository.Create(booking)
	if err != nil {
		c.JSON(500, ErrorResponse{
			Message: "Unable to create booking. \nError: "+ err.Error()})
		return
	}

	response := CreateBookingResponse{
		BookingId: bookingConfirm.BookingId,
		Total: float64(bookingConfirm.Total),
	}

	c.JSON(200, response)

}

func (h *Handler) FindAll(c *gin.Context){
	emailOrPhone := c.Query("userIdentifier")
	if emailOrPhone == "" {
		c.JSON(400, ErrorResponse{Message: "Invalid userIdentifier. Error: "})
		return
	}

	bookings, err := h.repository.FindAll(emailOrPhone)
	if err != nil {
		c.JSON(500, ErrorResponse{
			Message: "Unable to find booking. \nError: "+ err.Error()})
		return
	}

	c.JSON(200, bookings)
}