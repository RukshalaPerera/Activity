package routes

import (
	"JSONAPI/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(app *fiber.App) {
	app.Get("/", handler.GetAllProducts)
}
