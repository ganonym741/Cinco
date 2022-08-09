package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type AccountRepositoryInterface interface {
	FindById(ctx fiber.Ctx, cashflow model.Account) error
	Create(ctx fiber.Ctx, cashflow model.Account) error
}
