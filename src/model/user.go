package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email   string  `gorm:"unique;not null" json:"email" binding:"required"`
	Phone   int  `gorm:"unique;not null" json:"phone" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Password   string  `json:"password" binding:"required"`	
}