package helper

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(&ResponseError{
		StatusCode: status,
		Message:    err.Error(),
	})
}

func SendErrorMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(&ResponseError{
		StatusCode: status,
		Message:    message,
	})
}
