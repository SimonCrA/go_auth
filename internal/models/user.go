package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `grom:"uniqueIndex;not null" validate:"required,min=3,max=20"`
	Email    string `grom:"uniqueIndex;not null" validate:"required,email"`
	Password string `grom:"not null" validate:"required,min=6"`
	Lastname string `grom:"size:255;not null" validate:"required,min=3"`
	Name     string `grom:"size:255;not null" validate:"required,min=3"`
	IsActive bool   `grom:"default:false" validate:"boolean"`
	IsAdmin  bool   `grom:"default:false" validate:"boolean"`
}
