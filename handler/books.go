package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"books-data-scrapping/model"
	"books-data-scrapping/service"
)

type bookHandler struct {
	bookService service.IBookService
}

type IBookHandler interface {
	GetBooks(*fiber.Ctx) error
}

func NewBookHandler(service service.IBookService) bookHandler {
	return bookHandler{bookService: service}
}

func (handler bookHandler) GetBooks(c *fiber.Ctx) error {
	request := model.GetBooksRequest{
		Name: c.Query("name"),
	}

	request.Name = c.Query("name")

	if request.Name == "" {
		return c.Status(http.StatusBadRequest).SendString("Bad parameters")
	}

	response, err := handler.bookService.GetBooks(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}
	return c.Status(http.StatusOK).JSON(response)
}
