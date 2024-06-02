package main

import (
	"Project2/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()

	// Middleware to log requests
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Incoming request: %s %s", c.Method(), c.OriginalURL())
		return c.Next()
	})

	// CORS middleware with configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// Routes
	routes.SetUpDocumentRoutes(app)
	routes.SetUpRoleRoutes(app)
	routes.SetUPAuthRoutes(app)
	routes.SetupUserRoutes(app)
	routes.SetUPAuthUserRoutes(app)
	routes.SetUPDashboardRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
