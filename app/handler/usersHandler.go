package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	db "gitlab.com/cinco/pkg/postgres"
)

type CincoUser interface {
	UserRegister()
	UserLogin()
	UserLogout()
	UserProfile()
}

// func UserRegister(c *fiber.Ctx) error {

// }
// func UserLogin(c *fiber.Ctx) error {

// }
// func UserLogout(c *fiber.Ctx) error {

// }
func UserProfile(c *fiber.Ctx) error {
	db := db.DB
	var user model.User

	db.Find(&user)

	if user.UserId == "" {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "User data not found", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": user})
}
