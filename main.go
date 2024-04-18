package main

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Listen(":" + port)
}
