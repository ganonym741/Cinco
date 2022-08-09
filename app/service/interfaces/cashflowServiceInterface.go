package interfaces

import "gitlab.com/cinco/app/model"

type CashflowServiceInterface interface {
	FindTransactionLog(userUUID string, startDate int64, endDate int64) []model.Cashflow
}
