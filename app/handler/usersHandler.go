package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/service"
)

type Handler struct {
	service service.Service
}

//<<<<<<< HEAD
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
//=======
//func UserRegister(c *fiber.Ctx) error {
//	db := db.DB
//>>>>>>> feature/registerLogin/v1/cinco
//
//	inputUser := new(param.User)
//	inputUser.Id = uuid.New().String()
//	err := c.BodyParser(inputUser)
//	if err != nil {
//		return c.Status(500).JSON(fiber.Map{
//			"status":  "error",
//			"message": "Review your input",
//			"data":    err,
//		})
//	}
//	inputUser.Password, _ = utilities.GeneratePassword(inputUser.Password)
//	fmt.Println(inputUser.Password)
//
//	db.Create(&inputUser)
//
//<<<<<<< HEAD
// }

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
		return ctx.Status(201).
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
	fmt.Println("jalan")

	data, err := h.service.UserLogin(ctx, &paramsLogin)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   data.Token,
	})

}

////func UserLogout(c *fiber.Ctx) error {
////
////}
