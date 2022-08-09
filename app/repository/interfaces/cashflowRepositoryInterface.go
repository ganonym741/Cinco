package interfaces

import (
	"gitlab.com/cinco/app/model"
)

type CashflowRepositoryInterface interface {
	//PostTransaction(ctx fiber.Ctx, cashflow model.Cashflow) (model.Cashflow, error)
	FindByAccount(userUUID string, startDate int, endDate int) []model.Cashflow
}
