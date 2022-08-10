package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowServiceInterface interface {
	AddTransaction(ctx *fiber.Ctx, body *model.Cashflow) error
	FindTransactionLog(userUUID string, tipe string, startDate int64, endDate int64) []model.Cashflow
}
