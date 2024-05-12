package routes

import (
	"invoiceApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	api := app.Group("/auth")
	api.Get("/", controllers.Auth)
	api.Get("/profile", controllers.Auth)
	api.Post("/register", controllers.Register)
}

