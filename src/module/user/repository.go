package user

import (
	"BookingService/src/model"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Find(user User) (model.User, error)
	FindViaEmailOrPhone(email string, phone int) (model.User, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) FindViaEmailOrPhone(email string, phone int) (model.User, error){
	var user model.User
	if res := r.client.Where(
		"email=? or phone=?", email, phone).Find(&user); res.Error != nil {
		return model.User{}, res.Error
	}

	if user.ID == 0 {
		return model.User{}, errors.New("invalid user")
	}

	return user, nil
}

func (r *repository) Find(userDetails User) (model.User, error){
	var user model.User
	if res := r.client.Where(
		"phone=? and name=?", userDetails.Phone, 
		userDetails.Name).Find(&userDetails); res.Error != nil {
		return model.User{}, errors.New("invalid user")
	}
	if user.ID == 0 {
		return model.User{}, errors.New("invalid user")
	}

	return user, nil
}
