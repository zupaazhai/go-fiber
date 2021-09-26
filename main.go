package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zupaazhai/go-fiber/book"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book", book.GetBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
