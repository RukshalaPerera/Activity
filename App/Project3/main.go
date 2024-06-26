package main

import (
	"Project3/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Incoming request: %s %s", c.Method(), c.OriginalURL())
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	routes.SetupReservationRoutes(app)
	routes.SetupBookRoutes(app)
	routes.SetUpReportsRoutes(app)

	routes.SetUpDashboardRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8083"))
}
