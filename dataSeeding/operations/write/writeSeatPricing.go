package operations

import (
	"BookingService/initializer"
	"BookingService/src/model"
	"fmt"
)

func WriteSeatPricing(seatPricings []model.SeatPricing) {
	for i, seatPricing := range seatPricings {
		var dbSeatPricing model.SeatPricing
		if initializer.Db.Where(
			"seat_class=?", seatPricing.SeatClass,
			).Find(&dbSeatPricing).RowsAffected == 1 {
			fmt.Println("Seat Pricing found : ", i)
			continue
		}
		initializer.Db.Create(&seatPricing)
	}

	fmt.Println("............Write Seat Pricing Successfully............")
}
