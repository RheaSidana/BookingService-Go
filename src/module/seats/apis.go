package seats

import (
	"BookingService/initializer"
	"BookingService/src/module/seatBooked"
	"BookingService/src/module/seatPricing"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	seatHandler := InitHandler(repository, 
		seatBooked.InitRepository(initializer.Db),
		seatPricing.InitRepository(initializer.Db),
	)

	r.GET("/seats", seatHandler.FindAllSeats)
	r.GET("/seats/:id", seatHandler.FindSeat)
}
