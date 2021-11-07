package service

import (
	"books-data-scrapping/core/data"
	"books-data-scrapping/model"
	"fmt"
)

// Get Books data
func GetBooks(request model.GetBooksRequest) (*model.GetBooksResponse, error) {
	books := make([]model.Book, 0)
	err := data.GetBooksData(request.Name, &books)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := model.GetBooksResponse{
		Data: books,
	}
	return &response, nil
}
