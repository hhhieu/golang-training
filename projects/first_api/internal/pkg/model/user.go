package model

import (
	"gorm.io/gorm"
)

// User define the user model
type User struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	HashPassword string
	Email        string `gorm:"type:varchar(100);unique_index"`
}
