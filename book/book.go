package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get book list")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Get single book")
}
