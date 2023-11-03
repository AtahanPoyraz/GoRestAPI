package handlers

import (
	"github.com/AtahanPoyraz/gorest/database"
	"github.com/AtahanPoyraz/gorest/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func HomePage_(c *fiber.Ctx) error {
	message := map[string]string{"message": "Hello Go!!"}
	return c.Status(200).JSON(message)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
