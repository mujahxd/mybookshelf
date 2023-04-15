package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `gorm:"not null;size:255"`
	Author   string `gorm:"not null;size:255"`
	Year     int    `gorm:"not null"`
	ImageURL string `gorm:"size:255"`
	UserID   uint
}
