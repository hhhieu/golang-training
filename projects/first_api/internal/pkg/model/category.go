package model

import (
	"gorm.io/gorm"
)

// Category define the category model
type Category struct {
	gorm.Model
	Name    string
	Content string
}
