package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type AccountRepositoryInterface interface {
	UpdateBalance(ctx *fiber.Ctx, params string, balance int) error
	Create(account model.Account) error
	GetBalance(ctx *fiber.Ctx, params string) (int, error)
}
