package main

import (
	"JSONAPI/app/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetUpProductRoutes(app)

	fmt.Println("Starting server on http://localhost:3000")
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
