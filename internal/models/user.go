package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `grom:"uniqueIndex;not null"`
	Email    string `grom:"uniqueIndex;not null"`
	Password string `grom:"not null"`
	Name     string `grom:"size:255;not null"`
	Lastname string `grom:"size:255;not null"`
	IsActive bool   `grom:"default:false"`
	IsAdmin  bool   `grom:"default:false"`
}
