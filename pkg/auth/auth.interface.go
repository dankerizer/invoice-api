package auth

import "github.com/gofiber/fiber/v2"
type AuthInterface interface{
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}
