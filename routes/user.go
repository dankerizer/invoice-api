package routes

import (
	"invoiceApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	api := app.Group("/user")
	api.Get("/", controllers.User)

}
