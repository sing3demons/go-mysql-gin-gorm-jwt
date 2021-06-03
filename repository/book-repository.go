package repository

import (
	"github.com/sing3demons/golanh-api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book entity.Book) entity.Book
	UpdateBook(book entity.Book) entity.Book
	DeleteBook(book entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookRepositort struct{ connection *gorm.DB }

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositort{connection: db}
}

func (db *bookRepositort) InsertBook(book entity.Book) entity.Book {
	db.connection.Save(&book).Preload("User").Find(&book)
	return book
}
func (db *bookRepositort) UpdateBook(book entity.Book) entity.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}
func (db *bookRepositort) DeleteBook(book entity.Book) {
	db.connection.Delete(&book)
}
func (db *bookRepositort) AllBook() []entity.Book {
	var book []entity.Book
	db.connection.Preload("User").Find(&book)
	return book
}
func (db *bookRepositort) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}
