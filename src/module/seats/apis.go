package seats

import (
	"BookingService/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	seatHandler := InitHandler(repository)

	r.GET("/seats", seatHandler.FindAllSeats)
	r.GET("/seats/:id", seatHandler.FindSeat)
}
