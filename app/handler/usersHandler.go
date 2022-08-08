package handler

import "github.com/gofiber/fiber/v2"

type CincoUser interface {
	UserRegister()
	UserLogin()
	UserLogout()
	UserProfile()
}

func UserRegister(c *fiber.Ctx) {

}
func UserLogin(c *fiber.Ctx) {

}
func UserLogout(c *fiber.Ctx) {

}
func UserProfile(c *fiber.Ctx) {

}
