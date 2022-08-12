package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/service/interfaces"
)

type CincoAccount interface {
	AccountActivation(c *fiber.Ctx) error
}

type AccountHandler struct {
	AccountService interfaces.AccountServiceInterface
	UserService    interfaces.UserServiceInterface
}

func (a AccountHandler) AccountActivation(ctx *fiber.Ctx) error {
	userUUID := ctx.Params("userId")

	if userUUID != "" {
		user := a.UserService.FindByID(userUUID)

		if len(user.Id) == 0 || user.Id == "" {
			return ctx.Render("notification", fiber.Map{
				"name":    user.Fullname,
				"message": "User not found",
				"closure": "please contact your system administrator. Good Luck.",
			})
		} else {
			if user.Status {
				return ctx.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "your account is already activated",
					"closure": "please contact your system administrator. Good Luck.",
				})
			}
			err := a.AccountService.CreateAccount(userUUID)

			if err != nil {
				return ctx.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "error in creating account",
					"closure": "please contact your system administrator, Good Luck.",
				})
			}

			user.Status = true

			err = a.UserService.Update(user)

			if err != nil {
				return ctx.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "failed to activate your account",
					"closure": "please contact system administrator, Good Luck.",
				})
			}

			return ctx.Render("notification", fiber.Map{
				"name":    user.Fullname,
				"message": "your account has been activated successfully.",
				"closure": "Good Luck! Hope it works.",
			})
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "bad url request", "data": nil})
}

func NewAccountHandler(accountService interfaces.AccountServiceInterface, serviceInterface interfaces.UserServiceInterface) CincoAccount {
	return &AccountHandler{
		AccountService: accountService,
		UserService:    serviceInterface,
	}
}
