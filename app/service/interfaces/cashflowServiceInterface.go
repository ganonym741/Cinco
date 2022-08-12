package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowServiceInterface interface {
	AddTransaction(ctx *fiber.Ctx, body model.Cashflow) error
	FindTransactionLog(userUUID string, tipe string, startDate int64, endDate int64) []model.Cashflow
	EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*model.ResoponseCashflow, error)
	DeleteCashflow(ctx *fiber.Ctx, cashflowid string, paramsIdAccount string) (*model.Cashflow, error)
}
