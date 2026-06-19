package handlers

import (
	"input-api/models"
	"input-api/services"

	"github.com/gofiber/fiber/v2"
)

func QRHandler(c *fiber.Ctx) error {
	var req models.MatrixRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if !services.IsRectangular(req.Matrix) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Matrix must be rectangular",
		})
	}

	return c.JSON(fiber.Map{
		"received": req.Matrix,
	})
}
