package main

import (
	"github.com/AtahanPoyraz/gorest/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.HomePage_)

	app.Post("/fact", handlers.CreateFact)
}
