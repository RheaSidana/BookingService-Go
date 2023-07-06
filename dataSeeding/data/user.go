package data

import "BookingService/src/model"

func UsersData() []model.User {
	user1 := model.User{
		Name:     "abc",
		Email:    "abc@example.com",
		Phone:    1238907563,
		Password: "Abc_3",
	}
	user2 := model.User{
		Name:     "abcd",
		Email:    "abcd@example.com",
		Phone:    1238907573,
		Password: "Abcd_3",
	}
	user3 := model.User{
		Name:     "xyz",
		Email:    "XYZ@example.com",
		Phone:    1238907567,
		Password: "Xyz_3",
	}
	user4 := model.User{
		Name:     "pqr",
		Email:    "pqr@example.com",
		Phone:    1238907763,
		Password: "pqr_3",
	}
	user5 := model.User{
		Name:     "pqrt",
		Email:    "pqrt@example.com",
		Phone:    1233907563,
		Password: "Pqrt_3",
	}

	usersData := []model.User{user1, user2, user3, user4, user5}

	return usersData
}
