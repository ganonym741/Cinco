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
	userUUID := c.Params("userId")

	if userUUID != "" {
		user := a.UserService.FindByID(userUUID)

		if len(user.Id) == 0 || user.Id == "" {
			return c.Render("notification", fiber.Map{
				"name":    user.Fullname,
				"message": "User not found",
				"closure": "please contact your system administrator. Good Luck.",
			})
			//return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "User not found"})
		} else {
			if user.Status == true {
				//utilities.SendMail(user.Email, "your account is already activated, please contact system administrator.")
				//return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusConflict, "message": "user already activated"})
				return c.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "your account is already activated",
					"closure": "please contact your system administrator. Good Luck.",
				})
			}
			err := a.AccountService.CreateAccount(userUUID)

			if err != nil {
				//utilities.SendMail(user.Email, "error in creating account, please contact system administrator.")
				//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
				return c.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "error in creating account",
					"closure": "please contact your system administrator, Good Luck.",
				})
			}

			user.Status = true

			err = a.UserService.Update(user)

			if err != nil {
				//utilities.SendMail(user.Email, "failed to activate your account, please contact system administrator.")
				//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
				return c.Render("notification", fiber.Map{
					"name":    user.Fullname,
					"message": "failed to activate your account",
					"closure": "please contact system administrator, Good Luck.",
				})
			}

			//utilities.SendMail(user.Email, "your account has been activated successfully.")
			//return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User data has been activated successfully."})
			return c.Render("notification", fiber.Map{
				"name":    user.Fullname,
				"message": "your account has been activated successfully.",
				"closure": "Good Luck! Hope it works.",
			})
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
