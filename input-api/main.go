package main

import (
	"log"
	"os"

	"input-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Go API funcionando",
		})
	})

	app.Post("/received-matrix", handlers.QRHandler)

	log.Fatal(app.Listen(":" + port))

}
