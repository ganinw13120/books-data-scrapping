package handler_test

import (
	"books-data-scrapping/handler"
	"books-data-scrapping/model"
	"books-data-scrapping/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {

	type testCase struct {
		name     string
		keyword  string
		expected model.GetBooksResponse
	}

	cases := []testCase{
		{name: "Test 01", keyword: "asd", expected: model.GetBooksResponse{
			Data: []model.Book{
				{Name: "Test", Author: "Gan Mongklakorn", ImageUrl: "www.google.com"},
			},
		}},
	}
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			service := service.NewBookServiceMock()
			service.On("GetBooks").Return(v.expected.Data, nil)

			handler := handler.NewBookHandler(service, nil)

			app := fiber.New()
			app.Get("/books/get", handler.GetBooks)

			req := httptest.NewRequest("GET", fmt.Sprintf("/books/get?name=%s", v.keyword), nil)

			//Act
			res, _ := app.Test(req)
			defer res.Body.Close()

			//Assert
			if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
				body, _ := io.ReadAll(res.Body)
				var result model.GetBooksResponse
				err := json.Unmarshal(body, &result)

				assert.ErrorIs(t, nil, err)

				assert.Equal(t, result, v.expected)

				// assert.Equal(t, strconv.Itoa(expected), string(body))
			}

		})
	}
}
