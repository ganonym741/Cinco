package handler

import (
	"fmt"

	utilities "gitlab.com/cinco/utils"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/service/interfaces"
)

type CincoUser interface {
	UserRegister(ctx *fiber.Ctx) error
	UserLogin(ctx *fiber.Ctx) error
	UserLogout(ctx *fiber.Ctx) error
	UserProfile(ctx *fiber.Ctx) error
}

type UserHandler struct {
	UserService interfaces.UserServiceInterface
}

func (h UserHandler) UserRegister(ctx *fiber.Ctx) error {
	var params *param.User
	err := ctx.BodyParser(&params)
	if err != nil {
		fmt.Println("error1")
		return ctx.Status(400).
			JSON(fiber.Map{
				"status":   "failed",
				"messages": "correct your input",
			})
	}

	errors := utilities.ValidateStruct(*params)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	data, err := h.UserService.UserRegister(ctx, params)
	if err != nil {
		if err.Error() == "exist" {
			return ctx.Status(500).
				JSON(fiber.Map{
					"status": "failed",
					"data":   "exist",
				})
		}

		fmt.Println("error2")
		return ctx.Status(500).
			JSON(fiber.Map{
				"status": "failed",
				"data":   nil,
			})
	}

	return ctx.Status(201).
		JSON(fiber.Map{
			"status":  "success",
			"message": "Register Success Check Your Email to Activated",
			"data":    data,
		})
}

func (h UserHandler) UserProfile(ctx *fiber.Ctx) error {
	params := ctx.Query("id")
	data, err := h.UserService.GetUserDetail(ctx, params)
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

func (h UserHandler) UserLogin(ctx *fiber.Ctx) error {
	var paramsLogin param.Login
	err := ctx.BodyParser(&paramsLogin)
	if err != nil {
		return err
	}

	data, err := h.UserService.UserLogin(ctx, &paramsLogin)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   data.Token,
	})
}

func (h UserHandler) UserLogout(ctx *fiber.Ctx) error {
	params := ctx.Query("id")

	res, err := h.UserService.UserLogout(ctx, params)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func NewUserHandler(service interfaces.UserServiceInterface) CincoUser {
	return &UserHandler{
		UserService: service,
	}
}
