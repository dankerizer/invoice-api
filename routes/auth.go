package routes

import (
	"invoiceApi/controllers"
	"invoiceApi/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	api := app.Group("/auth")
	 jwt := middleware.Protected()
	api.Get("/", controllers.Auth)
	api.Get("/profile",jwt, controllers.Profile)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
}

