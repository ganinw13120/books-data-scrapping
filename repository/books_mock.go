package repository

import (
	"books-data-scrapping/model"

	"github.com/stretchr/testify/mock"
)

type bookRepositoryMock struct {
	mock.Mock
}

func NewBookRepositoryMock() *bookRepositoryMock {
	return &bookRepositoryMock{}
}

func (b *bookRepositoryMock) GetBooksData(name string, books *[]model.Book) error {
	args := b.Called()
	*books = args.Get(0).([]model.Book)
	return args.Error(1)
}
