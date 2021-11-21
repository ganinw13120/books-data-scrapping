package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"books-data-scrapping/model"

	"books-data-scrapping/service"
)

type bookHandler struct {
	bookService service.IBookService
	redis       *redis.Client
}

type IBookHandler interface {
	GetBooks(*fiber.Ctx) error
}

func NewBookHandler(service service.IBookService, redis *redis.Client) bookHandler {
	return bookHandler{
		bookService: service,
		redis:       redis,
	}
}

func (handler bookHandler) GetBooks(c *fiber.Ctx) error {
	request := model.GetBooksRequest{
		Name: c.Query("name"),
	}

	request.Name = c.Query("name")

	if request.Name == "" {
		return c.Status(http.StatusBadRequest).SendString("Bad parameters")
	}

	redisKey := request.Name
	if responseRedis, err := handler.redis.Get(context.Background(), redisKey).Result(); err == nil {
		var result *model.GetBooksResponse
		json.Unmarshal([]byte(responseRedis), &result)
		c.Set("Content-Type", "application/json")
		return c.Status(http.StatusOK).JSON(result)
	}

	response, err := handler.bookService.GetBooks(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}

	if data, err := json.Marshal(response); err == nil {
		handler.redis.Set(context.Background(), redisKey, string(data), 0)
	}

	return c.Status(http.StatusOK).JSON(response)
}
