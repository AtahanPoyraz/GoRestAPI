package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HomePage() {
	home := fiber.New()

	home.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go!")
	})

	home.Listen(":8080")
}