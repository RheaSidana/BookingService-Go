package main

import (
	"BookingService/initializer"
	"BookingService/src/module/booking"
	"BookingService/src/module/seats"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	seats.Apis(r)
	booking.Apis(r)

	r.Run()
}
