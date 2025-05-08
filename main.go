package main

import (
	"log"

	"test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
