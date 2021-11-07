package router

import (
	"books-data-scrapping/handler"

	"github.com/gofiber/fiber/v2"
)

type route struct {
	app *fiber.App
}

var router *route = nil

func New(app *fiber.App) *route {
	if router == nil {
		router = &route{app: app}
		router.setUp()
	}
	return router
}
func (r route) setUp() {
	group := r.app.Group("books")
	{
		group.Get("/get", handler.GetBooks)
	}
}
