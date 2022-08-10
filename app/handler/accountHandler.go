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

func (a AccountHandler) AccountActivation(c *fiber.Ctx) error {
	userUUID := c.Query("*")

	if userUUID != "" {
		user := a.UserService.FindByID(userUUID)

		if len(user.Id) == 0 || user.Id == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "User not found"})
		} else {
			err := a.AccountService.CreateAccount(userUUID)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
			}

			user.Status = true

			err = a.UserService.Update(user)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User data has been activated successfully."})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "bad url request", "data": nil})
}

func NewAccountHandler(accountService interfaces.AccountServiceInterface, serviceInterface interfaces.UserServiceInterface) CincoAccount {
	return &AccountHandler{
		AccountService: accountService,
		UserService:    serviceInterface,
	}
}
