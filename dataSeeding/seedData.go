package main

import (
	"BookingService/dataSeeding/data"
	r "BookingService/dataSeeding/operations"
	op "BookingService/dataSeeding/operations/create"
	w "BookingService/dataSeeding/operations/write"
	"BookingService/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	// \BookingService\dataSeeding\data\SeatPricing - MOCK_DATA (2).csv
	w.WriteSeatPricing(
		op.CreateSeatPricing(
			r.ReadCSV(`dataSeeding/data/SeatPricing - MOCK_DATA (2).csv`),
			),
		)
	
	// \BookingService\dataSeeding\data\Seats - MOCK_DATA (3).csv
	w.WriteSeat(
		op.CreateSeat(
			r.ReadCSV(`dataSeeding/data/Seats - MOCK_DATA (3).csv`),
			),
		)

	w.WriteUser(data.UsersData())
}
