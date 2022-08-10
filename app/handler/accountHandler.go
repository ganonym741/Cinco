package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/service/interfaces"
	utilities "gitlab.com/cinco/utils"
)

type CincoAccount interface {
	AccountActivation(c *fiber.Ctx) error
}

type AccountHandler struct {
	AccountService interfaces.AccountServiceInterface
	UserService    interfaces.UserServiceInterface
}

func (a AccountHandler) AccountActivation(ctx *fiber.Ctx) error {
	fmt.Printf("TEST %s ", ctx.Params("userId"))
	userUUID := ctx.Params("userId")

	if userUUID != "" {
		user := a.UserService.FindByID(userUUID)

		if len(user.Id) == 0 || user.Id == "" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "User not found"})
		} else {
			if user.Status {
				utilities.SendMail(user.Email, "your account is already activated, please contact system administrator.")
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusConflict, "message": "user already activated"})
			}
			err := a.AccountService.CreateAccount(userUUID)

			if err != nil {
				utilities.SendMail(user.Email, "error in creating account, please contact system administrator.")
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
			}

			user.Status = true

			err = a.UserService.Update(user)

			if err != nil {
				utilities.SendMail(user.Email, "failed to activate your account, please contact system administrator.")

				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
			}

			utilities.SendMail(user.Email, "your account has been activated successfully.")

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User data has been activated successfully."})
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
