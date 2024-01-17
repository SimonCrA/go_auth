package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" validate:"required,min=3,max=12"`
	Email    string `gorm:"uniqueIndex;not null" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Lastname string `gorm:"size:255;not null" validate:"required,min=3"`
	Name     string `gorm:"size:255;not null" validate:"required,min=3"`
	IsActive bool   `gorm:"default:false" validate:"boolean"`
	IsAdmin  bool   `gorm:"default:false" validate:"boolean"`
}
