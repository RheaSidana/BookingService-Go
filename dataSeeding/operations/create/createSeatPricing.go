package operations

import (
	"BookingService/src/model"
	"fmt"
	"strconv"
	"strings"
)

// func printSeatPricing(seatPricing model.SeatPricing){
// 	fmt.Println("Seat Pricing :")
// 	fmt.Println("Seat Class : ", seatPricing.SeatClass)
// 	fmt.Println("Min Price : ", seatPricing.MinPrice)
// 	fmt.Println("Max Price : ", seatPricing.MaxPrice)
// 	fmt.Println("Normal Price : ", seatPricing.NormalPrice)
// 	fmt.Println()
// }

func mapToSeatPricing(index int, value string, seatPricing model.SeatPricing) model.SeatPricing {
	switch index {
	case 1:
		seatPricing.SeatClass = value
	case 2:
		value = strings.TrimLeft(value, "$")
		seatPricing.MinPrice, _ = strconv.ParseFloat(value, 64)
	case 3:
		value = strings.TrimLeft(value, "$")
		seatPricing.NormalPrice, _ = strconv.ParseFloat(value, 64)
	case 4:
		value = strings.TrimLeft(value, "$")
		seatPricing.MaxPrice, _ = strconv.ParseFloat(value, 64)
	}

	return seatPricing
}

func createSeatPricing(record []string) model.SeatPricing {

	var seatPricing model.SeatPricing

	for i, field := range record {
		seatPricing = mapToSeatPricing(i, string(field), seatPricing)
	}

	return seatPricing
}

func CreateSeatPricing(records [][]string) []model.SeatPricing {
	var seatPricings []model.SeatPricing

	for i, record := range records {
		if i == 0 {
			continue
		}
		seatPricing := createSeatPricing(record)
		// printSeatPricing(seatPricing)
		seatPricings = append(seatPricings, seatPricing)
	}

	fmt.Println("............Create Seat Pricing Successfully............")

	return seatPricings
}
