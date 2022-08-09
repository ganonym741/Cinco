package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type AccountRepositoryInterface interface {
	Create(ctx fiber.Ctx, account model.Account) error
}
