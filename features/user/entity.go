package user

import (
	"github.com/mujahxd/mybookshelf/features/book"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"not null;size:255"`
	Email          string `gorm:"not null;size:255;uniqueIndex"`
	PasswordHash   string `gorm:"not null"`
	AvatarFileName string `gorm:"size:255"`
	Books          []book.Book
}
