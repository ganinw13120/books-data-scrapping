package main

import (
	"books-data-scrapping/router"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
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

func setupFiber() error {
	app := fiber.New(fiberConfig())

	router.New(app)
	err := app.Listen(":" + os.Getenv("PORT"))

	return err
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
