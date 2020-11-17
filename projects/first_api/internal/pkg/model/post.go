package model

import (
	"time"

	"gorm.io/gorm"
)

// Post define the model of blog post
type Post struct {
	gorm.Model
	AuthorID   int
	User       User `gorm:"foreignKey:AuthorID"`
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	Title      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Content    string
	Rating     int `gorm:"size:5"`
}
