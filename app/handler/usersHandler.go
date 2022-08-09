package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{
		service: s,
	}
}

// type CincoUser interface {
// 	UserRegister()
// 	UserLogin()
// 	UserLogout()
// 	UserProfile()
// }

// func UserRegister(c *fiber.Ctx) error {

// }
// func UserLogin(c *fiber.Ctx) error {

// }
// func UserLogout(c *fiber.Ctx) error {

// }
func (h Handler) UserProfile(c *fiber.Ctx) error {
	ctx := context.Background()
	params := c.Params("userId")
	data, err := h.service.GetUserDetail(ctx, params)
	if err != nil {
		return c.Status(201).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}

	return c.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}
