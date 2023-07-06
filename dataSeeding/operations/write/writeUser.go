package operations

import (
	"BookingService/src/model"
	"BookingService/initializer"
)

func WriteUser(usersData []model.User) {
	for _, user := range usersData {
		if initializer.Db.Where("email=?", user.Email).Find(&user).RowsAffected == 1 {
			continue
		}
		initializer.Db.Create(&user)
	}
}