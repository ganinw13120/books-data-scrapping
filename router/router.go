package router

import (
	"github.com/gofiber/fiber/v2"

	"books-data-scrapping/handler"
)

type route struct {
	app *fiber.App

	bookHandler handler.IBookHandler
}

var router *route = nil

func New(app *fiber.App, bookHandler handler.IBookHandler) *route {
	if router == nil {
		router = &route{
			app:         app,
			bookHandler: bookHandler,
		}
		router.setUp()
	}
	return router
}
func (r route) setUp() {
	group := r.app.Group("books")
	{
		group.Get("/get", r.bookHandler.GetBooks)
	}
}
