package service

import (
	"errors"
	"fmt"
	"time"

	utilities "gitlab.com/cinco/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type Service struct {
	cashflowRepository interfaces.CashflowRepositoryInterface
	account            interfaces.AccountRepositoryInterface
}

func (s Service) AddTransaction(ctx *fiber.Ctx, body model.Cashflow) error {
	body.Id = uuid.NewString()

	balance, err := s.account.GetBalance(ctx, body.AccountId)
	if err != nil {
		return err
	}

	fmt.Println("cek staus bod")

	if body.Type == "debet" {
		balance = balance + body.Amount
	} else if body.Type == "credit" {
		if balance > body.Amount {
			balance = balance - body.Amount
		} else {
			return nil // saldo tidak mencukupi
		}
	} else {
		return nil // tipe transaksi anda tidak valid
	}

	body.BalanceHistory = balance

	err = s.account.UpdateBalance(ctx, body.AccountId, balance)
	if err != nil {
		fmt.Println("cek staus bod 4")
		return err
	}

	err = s.cashflowRepository.PostTransaction(ctx, &body)
	if err != nil {
		fmt.Println("cek staus bod 5")
		return err
	}
	return nil
}

func (s Service) FindTransactionLog(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error) {
	return s.cashflowRepository.FindByAccount(userUUID, tipe, utilities.Bod(startDate), utilities.Eod(endDate))
}

func (s Service) EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*model.ResoponseCashflow, error) {

	data := model.Cashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}

	// 	fmt.Println("ini data input", data.Amount)

	amountnhistory, amounttypes, _, err := s.cashflowRepository.GetHistoryandAmountBefore(ctx, params)
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
			balance = balance - (data.Amount - amountnhistory)
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
	return &model.ResoponseCashflow{
		Description:    body.Description,
		Amount:         body.Amount,
		BalanceHistory: balance,
	}, nil
}

func (s Service) DeleteCashflow(ctx *fiber.Ctx, cashflowid string, paramsIdAccount string) error {
	amountHistory, cashflowTypes, _, err := s.cashflowRepository.GetHistoryandAmountBefore(ctx, cashflowid)
	if err != nil {
		return err
	}

	if cashflowTypes == "credit" {
		amountHistory = -amountHistory
	} else if cashflowTypes == "debet" {

	} else {
		return errors.New("tipe transaksi tidak bisa terbaca")
	}

	err = s.cashflowRepository.RepoUpdateBalance(ctx, amountHistory, paramsIdAccount)
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
