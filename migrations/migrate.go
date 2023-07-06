package main

import (
	"BookingService/initializer"
	"BookingService/src/model"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.Db.AutoMigrate(&model.SeatPricing{})
	initializer.Db.AutoMigrate(&model.Seat{})
	initializer.Db.AutoMigrate(&model.User{})
	initializer.Db.AutoMigrate(&model.Booking{})
	initializer.Db.AutoMigrate(&model.SeatBooked{})
}
