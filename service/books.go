package service

import (
	"books-data-scrapping/model"
	"books-data-scrapping/repository"
	"fmt"
)

type bookService struct {
	bookRepository repository.IBookRepository
}

type IBookService interface {
	GetBooks(name string) (*model.GetBooksResponse, error)
}

func NewBookService(bookRepository repository.IBookRepository) bookService {
	return bookService{bookRepository: bookRepository}
}

// Get Books data
func (service bookService) GetBooks(name string) (*model.GetBooksResponse, error) {
	books := make([]model.Book, 0)

	err := service.bookRepository.GetBooksData(name, &books)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := model.GetBooksResponse{
		Data: books,
	}
	return &response, nil
}
