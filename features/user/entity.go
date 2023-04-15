package user

import (
	"github.com/mujahxd/mybookshelf/features/book"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"not null;size:255"`
	Email          string `gorm:"not null;size:255;uniqueIndex"`
	Password       string `gorm:"not null;size:255"`
	AvatarFileName string `gorm:"size:255"`
	Books          []book.Book
}

// Name           string            `json:"name" gorm:"type:varchar(45);not null"`
// 	Username       string            `json:"username" gorm:"type:varchar(12);unique;primaryKey"`
// 	Password       string            `json:"password" gorm:"type:varchar(255);not null"`
// 	AvatarFileName string            `json:"avatar" gorm:"type:varchar(255)"`
