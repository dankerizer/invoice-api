package main

import (
	"encoding/json"
	"invoiceApi/database"
	"invoiceApi/routes"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)
func init() {
    if err := godotenv.Load(); err != nil {
        log.Panic("No .env file found")
    }
}
func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			return
		}
	}()

	database.ConnectDb()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		Immutable:         true,
		BodyLimit:         50 * 1024 * 1024,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Invoice.id")
	})

	app.Use(func(c *fiber.Ctx) error {
		//c.Locals("body", string(c.Body()))
		payload := map[string]interface{}{}
		if c.Method() == "POST" || c.Method() == "PATCH" {
			err := json.Unmarshal(c.Body(), &payload)
			if err != nil {
				log.Error(err)
			}
		}

		// convert JSON into one-line string
		oneLinePayload, _ := json.Marshal(payload)
		c.Locals("body", strings.ReplaceAll(string(oneLinePayload), "\n", "\\n"))
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} Body: ${locals:body}\n",
		Output: os.Stdout,
	}))

	app.Use(cors.New(cors.Config{
		// AllowCredentials: true,
		AllowOrigins: "*",
	}))
	routes.SetupRouter(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("@dankedev: Hayo, cari apa nih? ")
	})

	app.Listen(":8000")
}
