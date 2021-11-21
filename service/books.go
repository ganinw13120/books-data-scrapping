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
	GetBooks(request model.GetBooksRequest) (*model.GetBooksResponse, error)
}

func NewBookController(bookRepository repository.IBookRepository) bookService {
	return bookService{bookRepository: bookRepository}
}

// Get Books data
func (service bookService) GetBooks(request model.GetBooksRequest) (*model.GetBooksResponse, error) {
	books := make([]model.Book, 0)

	err := service.bookRepository.GetBooksData(request.Name, &books)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := model.GetBooksResponse{
		Data: books,
	}
	return &response, nil
}
