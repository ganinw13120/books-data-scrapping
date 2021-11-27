package main

import (
	"books-data-scrapping/handler"
	"books-data-scrapping/repository"
	"books-data-scrapping/router"
	"books-data-scrapping/service"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func setupTimeZone() error {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}
	time.Local = location
	return nil
}

func fiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Books",
	}
}

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}
}

func setupFiber() error {
	app := fiber.New(fiberConfig())
	app.Use(cors.New(corsConfig()))
	app.Use(recover.New())
	redis := setupRedis()

	bookRepository := repository.NewBookRepository()

	bookService := service.NewBookController(bookRepository)

	bookHandler := handler.NewBookHandler(bookService, redis)

	router.New(app, bookHandler)
	err := app.Listen(":" + os.Getenv("PORT"))

	return err
}

func setupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = setupTimeZone()
	if err != nil {
		panic(err)
	}
	err = setupFiber()
	if err != nil {
		panic(err)
	}
}
