package model

// User define the user model
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null" form:"name" json:"name"`
	Email string `gorm:"type:varchar(100);unique_index" form:"email" json:"email"`
}
