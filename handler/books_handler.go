package handler

import (
	"books-data-scrapping/core/service"
	"books-data-scrapping/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	request := model.GetBooksRequest{}

	request.Name = c.Query("name")

	if request.Name == "" {
		return c.Status(http.StatusBadRequest).SendString("Bad parameters")
	}

	response, err := service.GetBooks(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}
	return c.Status(http.StatusOK).JSON(response)
}
