package service

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
	"time"
)

type CashflowService struct {
	cashflowRepository interfaces.CashflowRepositoryInterface
	account            interfaces.AccountRepositoryInterface
}

func (c CashflowService) FindTransactionLog(userUUID string, tipe string, startDate int64, endDate int64) []model.Cashflow {
	return c.cashflowRepository.FindByAccount(userUUID, tipe, time.Unix(startDate, 0), time.Unix(endDate, 0))
}

/*func (s Service) AddTransaction(ctx context.Context, userid string) (*model.User, error) {
	var data model.User
	err := s.repository.GetUserDetail(ctx, &data, userid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}*/

func NewCashflowService(repository interfaces.CashflowRepositoryInterface) serviceInterface.CashflowServiceInterface {
	return &CashflowService{
		cashflowRepository: repository,
	}
}
