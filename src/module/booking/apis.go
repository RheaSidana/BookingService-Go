package booking

import (
	"BookingService/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	bookingHandler := InitHandler(repository)

	r.POST("/booking", bookingHandler.Create)
	r.GET("/bookings", bookingHandler.FindAll)
}
