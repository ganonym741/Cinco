package interfaces

import (
	"gitlab.com/cinco/app/model"
	"time"
)

type CashflowRepositoryInterface interface {
	//PostTransaction(ctx fiber.Ctx, cashflow model.Cashflow) (model.Cashflow, error)
	FindByAccount(userUUID string, startDate time.Time, endDate time.Time) []model.Cashflow
}
