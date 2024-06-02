package routes

import (
	"Project3/app/dashboard"
	"Project3/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(app *fiber.App) {
	app.Get("/books", handler.GetAllBooks)
	app.Get("/book/:_id", handler.GetABook)
	app.Post("/books", handler.CreateBook)
	app.Put("/books/:_id", handler.EditABook)
	app.Delete("/books/:_id", handler.DeleteABook)
	app.Get("/books/search", handler.SearchBooks)
	app.Get("/books", handler.GetBookList)
	app.Get("/Books/count", dashboard.BookCount)
}

func SetupReservationRoutes(app *fiber.App) {
	app.Get("/reservations", handler.GetAllReservations)
	app.Get("/reservation/:id", handler.GetAReservation)
	app.Post("/reservations", handler.CreateAReservation)
	app.Put("/reservations/:id", handler.EditAReservation)
	app.Delete("/reservations/:id", handler.DeleteAReservation)
}

func SetUpReportsRoutes(app *fiber.App) {
	app.Get("/generate-reservation-report", handler.GenerateAllReservationReportHandler)
}

func SetUpDashboardRoutes(app *fiber.App) {
	app.Get("/Books/Chart1", dashboard.GenerateMonthlyBookChart)
}
