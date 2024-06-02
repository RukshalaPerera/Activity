package routes

import (
	"Project2/app/dashboard"
	"Project2/app/handler"
	"Project2/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// user
func SetupUserRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAllUsers)
	app.Get("/user/:_id", handler.GetAUser)
	app.Post("/users", handler.CreateUser)
	app.Put("/user/:_id", handler.EditAUser)
	app.Delete("/user/:_id", handler.DeleteAUser)
	app.Get("/users/search", handler.SearchUsers)
}

// role
func SetUpRoleRoutes(app *fiber.App) {
	app.Get("/roles", handler.GetAllRoles)
	app.Get("/role/:_id", handler.GetARole)
	app.Post("/roles", handler.CreateRole)
	app.Put("role/:_id", handler.EditARole)
	app.Delete("/role/:_id", handler.DeleteARole)
}

// document
func SetUpDocumentRoutes(app *fiber.App) {
	app.Get("/documents", middleware.DocumentAuthorization("user", "admin", "moderator"), handler.ListDocuments)       //list documents
	app.Get("/document/:id", middleware.DocumentAuthorization("user", "admin", "moderator"), handler.DownloadDocument) //downloading by ID
	app.Post("/documents", middleware.DocumentAuthorization("moderator"), handler.UploadDocument)                      //filename and title
}

// auth
func SetUPAuthRoutes(app *fiber.App) {
	app.Post("/login", handler.Login)
	app.Post("/signup", handler.SignUp)
	app.Post("/home", handler.Home)
	app.Get("/premium", handler.Premium)
	app.Get("/logout", handler.Logout)
}

// AuthUser
func SetUPAuthUserRoutes(app *fiber.App) {
	app.Post("/AuthUsers", handler.CreateAuthUser)
	app.Get("/AuthUsers", handler.GetAllAuthUsers)
	app.Get("/AuthUser/:_id", handler.GetAuthUser)
	app.Put("/AuthUser/:_id", handler.EditAuthUser)
	app.Delete("/AuthUser/:_id", handler.DeleteAuthUser)
}

// dashboard
func SetUPDashboardRoutes(app *fiber.App) {
	app.Get("/AuthUsers/count", dashboard.AuthUserCount)
	app.Get("/Users/count", dashboard.UserCount)
}
