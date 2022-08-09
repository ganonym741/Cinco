package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowRepositoryInterface interface {
	//PostTransaction(ctx fiber.Ctx, cashflow model.Cashflow) (model.Cashflow, error)
	FindByAccount(ctx fiber.Ctx, userUUID string, startDate int, endDate int) []model.Transaction_log
}
