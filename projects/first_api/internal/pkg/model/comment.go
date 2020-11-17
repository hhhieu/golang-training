package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment define the comment model
type Comment struct {
	gorm.Model
	AuthorID  int
	User      User `gorm:"foreignKey:AuthorID"`
	PostID    int
	Post      Post `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}
