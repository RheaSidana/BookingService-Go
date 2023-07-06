package operations

import (
	"BookingService/initializer"
	"BookingService/src/model"
	"fmt"
)

func WriteSeat(seats []model.Seat) {
	for i, seat := range seats {
		var dbSeat model.Seat
		if initializer.Db.Where(
			"seat_identifier=?", seat.SeatIdentifier,
			).Find(&dbSeat).RowsAffected == 1 {
			fmt.Println("Seat found : ", i)
			continue
		}
		initializer.Db.Create(&seat)
	}

	fmt.Println("............Write Seat Successfully............")
}
