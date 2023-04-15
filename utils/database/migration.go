package database

import (
	"github.com/mujahxd/mybookshelf/features/book"
	"github.com/mujahxd/mybookshelf/features/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users := user.User{}
	books := book.Book{}

	db.AutoMigrate(users)
	db.AutoMigrate(books)
}
