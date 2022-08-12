package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/service"
	utilities "gitlab.com/cinco/utils"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) UserRegister(ctx *fiber.Ctx) error {
	var params *param.User
	err := ctx.BodyParser(params)
	if err != nil {
		return ctx.Status(400).
			JSON(fiber.Map{
				"status": "failed",
				"data":   nil,
			})
	}

	errors := utilities.ValidateStruct(*params)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	data, err := h.service.UserRegister(ctx, params)
	if err != nil {
		return ctx.Status(400).
			JSON(fiber.Map{
				"status": "failed",
				"data":   nil,
			})
	}

	return ctx.Status(201).
		JSON(fiber.Map{
			"status":  "success",
			"message": "User data created",
			"data":    data,
		})
}

func (h Handler) UserProfile(ctx *fiber.Ctx) error {
	params := ctx.Query("id")
	data, err := h.service.GetUserDetail(ctx, params)
	if err != nil {
		return ctx.Status(404).
			JSON(fiber.Map{
				"status":  "failed",
				"message": "Data not found",
				"data":    nil,
			})
	}
	return ctx.Status(201).
		JSON(fiber.Map{
			"status":  "success",
			"message": "User data retrieved",
			"data":    data,
		})
}

func (h Handler) UserLogin(ctx *fiber.Ctx) error {
	var paramsLogin param.Login
	err := ctx.BodyParser(&paramsLogin)
	if err != nil {
		return err
	}

	data, err := h.service.UserLogin(ctx, &paramsLogin)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   data.Token,
	})

}

func (h Handler) UserLogout(ctx *fiber.Ctx) error {
	params := ctx.Query("id")

	res, err := h.service.UserLogout(ctx, params)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}
