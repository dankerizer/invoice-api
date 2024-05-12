package controllers

import (
	"invoiceApi/database"
	"invoiceApi/helper"
	"invoiceApi/middleware"
	"invoiceApi/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c *fiber.Ctx) error {
	return c.SendString("Hello Hadie ini adalah router auth!!")
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, err)

	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	} // If the email is present in the DB then compare the Passwords and if incorrect password then return error.

	token, err := middleware.GenerateJWTToken(&user)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// cookie := fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token.Token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }

	return c.Status(fiber.StatusOK).JSON(token)

	// c.Cookie(&cookie)

	// return c.JSON(fiber.Map{
	// 	"message": "success",
	// 	"token":   token.Token,
	// 	"expired": token.Expired,
	// })
}
func Logout(c *fiber.Ctx) error {
	return c.SendString("Hello Hadie ini adalah router logout!!")

}
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e. (14 in our case).

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.DB.Create(&user) //Adds the data to the DB
	return c.JSON(data)
}

func Profile(c *fiber.Ctx) error {
	// cookie := c.Cookies("jwt")

	token := c.Locals("user");
			// claims := token.Claims

	return c.Status(fiber.StatusOK).JSON(token)
}
