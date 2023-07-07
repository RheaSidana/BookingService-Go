package seats

import (
	"BookingService/src/module/seatBooked"
	"BookingService/src/module/seatPricing"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
	seatBookedRepo seatBooked.Repository
	seatPricingRepo seatPricing.Repository
}

func (h *Handler) FindAllSeats(c *gin.Context){
	seats, err := h.repository.FindAll() 
	if err != nil {
		c.JSON(500, ErrorResponse{
			Message: "Unable to find seats. Error: " + err.Error()})
		return
	}

	seatsBooked, err := h.seatBookedRepo.FindAll()
	if err != nil {
		c.JSON(500, ErrorResponse{
			Message: "Unable to find seats. Error: " + err.Error()})
		return
	}

	seatsList := createSeatsList(seats, seatsBooked)

	c.JSON(200, seatsList)
}

func (h *Handler) FindSeat(c *gin.Context) {
	seatId := c.Param("id")
	if seatId == "" {
		c.JSON(400, ErrorResponse{Message: "Invalid ID. Error: "})
		return
	}

	seat, err := h.repository.FindSeat(seatId)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Seat not Found. Error: " + err.Error()})
		return
	}

	seatPrice, err := h.seatPricingRepo.Find(seat.SeatClass)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Seat not Found. Error: " + err.Error()})
		return
	}

	countR, err := h.repository.Count(seatPrice.SeatClass)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Seat not Found. Error: " + err.Error()})
		return
	}

	countB, err := h.seatBookedRepo.Count(seat.SeatClass)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Seat not Found. Error: " + err.Error()})
		return
	}

	temp := Temp{
		Seat: seat,
		SeatPrice: seatPrice,
		TotalRows: countR,
		TotalBooked: countB,
	}

	seatWithPrice := createSeatIdPrice(temp)

	c.JSON(200, seatWithPrice)
}