package service

import (
	"errors"
	"time"

	utilities "gitlab.com/cinco/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gitlab.com/cinco/app/response"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type Service struct {
	cashflowRepository interfaces.CashflowRepositoryInterface
	account            interfaces.AccountRepositoryInterface
}

func (s Service) TotalCashflow(userUUID string, startDate time.Time, endDate time.Time) (model.Total, error) {
	return s.cashflowRepository.FindTotal(userUUID, utilities.Bod(startDate), utilities.Eod(endDate))
}

func (s Service) AddTransaction(ctx *fiber.Ctx, body model.Cashflow) error {
	body.Id = uuid.NewString()

	balance, err := s.account.GetBalance(ctx, body.AccountId)
	if err != nil {
		return err
	}

	if body.Type == "debet" {
		balance = balance + body.Amount
	} else if body.Type == "credit" {
		if balance > body.Amount {
			balance = balance - body.Amount
		} else {
			return errors.New("saldo tidak mencukupi")
		}
	}

	err = s.account.UpdateBalance(ctx, body.AccountId, balance)
	if err != nil {
		return err
	}

	err = s.cashflowRepository.PostTransaction(ctx, &body)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) FindTransactionLog(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error) {
	return s.cashflowRepository.FindByAccount(userUUID, tipe, utilities.Bod(startDate), utilities.Eod(endDate))
}

func (s Service) EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*response.ResponseCashflow, error) {

	data := model.Cashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}

	amountnhistory, amounttypes, err := s.cashflowRepository.GetHistoryandAmountBefore(ctx, params)
	if err != nil {
		return nil, err
	}

	balance, err := s.account.GetBalance(ctx, paramsIdAccount)
	if err != nil {
		return nil, err
	}

	switch amounttypes {
	case "credit":
		if data.Amount > amountnhistory {
			balance = balance - (data.Amount - amountnhistory)
		} else {
			balance = balance + (data.Amount - amountnhistory)
		}
	case "debet":
		if data.Amount > amountnhistory {
			balance = balance + (data.Amount - amountnhistory)
		} else {
			balance = balance - (amountnhistory - data.Amount)
		}
	}

	err = s.cashflowRepository.RepoUpdateBalance(ctx, balance, paramsIdAccount)
	if err != nil {
		return nil, err
	}
	err = s.cashflowRepository.RepoEditCashFlow(ctx, &data, params)
	if err != nil {
		return nil, err
	}

	return &response.ResponseCashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}, nil
}

func (s Service) DeleteCashflow(ctx *fiber.Ctx, cashflowid string, paramsIdAccount string) error {
	amountHistory, cashflowTypes, err := s.cashflowRepository.GetHistoryandAmountBefore(ctx, cashflowid)
	if err != nil {
		return err
	}

	balance, err := s.account.GetBalance(ctx, paramsIdAccount)
	if err != nil {
		return err
	}

	if cashflowTypes == "credit" {
		balance = balance + amountHistory
	} else if cashflowTypes == "debet" {
		balance = balance - amountHistory
	} else {
		return errors.New("tipe transaksi tidak bisa terbaca")
	}

	err = s.cashflowRepository.RepoUpdateBalance(ctx, balance, paramsIdAccount)
	if err != nil {
		return err
	}

	err = s.cashflowRepository.DeleteCashflow(ctx, cashflowid)
	if err != nil {
		return err
	}

	return nil
}

func NewCashflowService(repository interfaces.CashflowRepositoryInterface, accountRepository interfaces.AccountRepositoryInterface) serviceInterface.CashflowServiceInterface {
	return &Service{
		cashflowRepository: repository,
		account:            accountRepository,
	}
}
