package service

import (
	"books-data-scrapping/model"

	"github.com/stretchr/testify/mock"
)

type bookServiceMock struct {
	mock.Mock
}

func NewBookServiceMock() *bookServiceMock {
	return &bookServiceMock{}
}

func (service bookServiceMock) GetBooks(name string) (*model.GetBooksResponse, error) {
	args := service.Called()
	result := &model.GetBooksResponse{
		Data: args.Get(0).([]model.Book),
	}
	return result, args.Error(1)
}
