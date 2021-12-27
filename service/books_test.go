package service_test

import (
	"books-data-scrapping/model"
	"books-data-scrapping/repository"
	"books-data-scrapping/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {

	type testCase struct {
		keyword  string
		expected *model.GetBooksResponse
	}

	cases := []testCase{
		{keyword: "Test", expected: &model.GetBooksResponse{
			Data: []model.Book{
				{Name: "Test", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
			},
		}},
		{keyword: "Multiple", expected: &model.GetBooksResponse{
			Data: []model.Book{
				{Name: "Test", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
				{Name: "ATOMIC HABITS เพราะชีวิตดีได้กว่าที่เป็น ", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
				{Name: "Atomic Habits", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
				{Name: "Atomic Bomb of Brand", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
			},
		}},
	}

	for _, v := range cases {
		t.Run(v.keyword, func(t *testing.T) {
			bookRepository := repository.NewBookRepositoryMock()

			bookRepository.On("GetBooksData").Return(v.expected.Data, nil)

			bookService := service.NewBookService(bookRepository)

			result, err := bookService.GetBooks(v.keyword)

			assert.Equal(t, v.expected, result)
			assert.ErrorIs(t, nil, err)
		})
	}

}
