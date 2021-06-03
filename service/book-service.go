package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/sing3demons/golanh-api/dto"
	"github.com/sing3demons/golanh-api/entity"
	"github.com/sing3demons/golanh-api/repository"
)

type BookService interface {
	Insert(book dto.BookCreateDTO) entity.Book
	Update(bookUpdate dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}

}

func (service *bookService) Insert(bookCreate dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	if err := smapping.FillStruct(&book, smapping.MapFields(&bookCreate)); err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	resp := service.bookRepository.InsertBook(book)
	return resp
}
func (service *bookService) Update(bookUpdate dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&bookUpdate))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(book entity.Book) {
	service.bookRepository.DeleteBook(book)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
