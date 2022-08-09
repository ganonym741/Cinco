package service

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type CashflowService struct {
	cashflowRepository interfaces.CashflowRepositoryInterface
}

func (c CashflowService) FindTransactionLog(userUUID string, startDate int, endDate int) []model.Cashflow {
	//TODO implement me
	panic("implement me")
}

func NewCashflowService(repository interfaces.CashflowRepositoryInterface) serviceInterface.CashflowServiceInterface {
	return &CashflowService{
		cashflowRepository: repository,
	}
}
