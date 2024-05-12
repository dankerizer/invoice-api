package controllers

import (
	"github.com/gofiber/fiber/v2"
)
func User(c *fiber.Ctx) error {

	return c.SendString("Hello Hadie ini adalah router user!!")
}
