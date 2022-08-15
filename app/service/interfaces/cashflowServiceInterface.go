package interfaces

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/response"
)

type CashflowServiceInterface interface {
	AddTransaction(ctx *fiber.Ctx, body model.Cashflow) error
	TotalCashflow(userUUID string, startDate time.Time, endDate time.Time) (model.Total, error)
	FindTransactionLog(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error)
	EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*response.ResponseCashflow, error)
	DeleteCashflow(ctx *fiber.Ctx, cashflowid string, paramsIdAccount string) error
}
