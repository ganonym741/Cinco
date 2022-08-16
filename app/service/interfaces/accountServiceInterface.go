package interfaces

import "github.com/gofiber/fiber/v2"

type AccountServiceInterface interface {
	CreateAccount(userUUID string) error
	GetBalance(ctx *fiber.Ctx, params string) (int, error)
}
