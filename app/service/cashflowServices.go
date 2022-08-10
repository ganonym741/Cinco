package service

import (
	"fmt"
	"time"

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

func (s Service) AddTransaction(ctx *fiber.Ctx, body *model.Cashflow) error {
	body.Id = uuid.NewString()

	balance, err := s.account.GetBalance(ctx, body.AccountId)
	if err != nil {
		return err
	}

	if body.Type == "debet" {
		balance = balance + body.Amount
	} else if body.Type == "kredit" {
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
		return err
	}

	err = s.cashflowRepository.PostTransaction(ctx, body)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) FindTransactionLog(userUUID string, tipe string, startDate int64, endDate int64) []model.Cashflow {
	return s.cashflowRepository.FindByAccount(userUUID, tipe, time.Unix(startDate, 0), time.Unix(endDate, 0))
}
func (s Service) DeleteCashflow(ctx *fiber.Ctx, cashflowid string) (*model.Cashflow, error) {
	var data model.Cashflow

	err := s.cashflowRepository.DeleteCashflow(ctx, &data, cashflowid)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (s Service) EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*model.Cashflow, error) {

	data := model.Cashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}

	fmt.Println("ini data input", data.Amount)

	amountnhistory, amounttypes, balancehistory, err := s.cashflowRepository.GetHistoryandAmountBefore(ctx, params)
	if err != nil {
		return nil, err
	}

	fmt.Println("S ini amount history", amountnhistory)

	balance, err := s.account.GetBalance(ctx, paramsIdAccount)
	if err != nil {
		return nil, err
	}

	fmt.Println("S ini balance dari repo getbalance", balance)

	accountUpdate := model.Account{}

	fmt.Println("S ini isi balance dari variable accountupdate ", accountUpdate.Balance)

	fmt.Println("S ini type", amountnhistory)
	switch amounttypes {
	case "credit":
		if balance > data.Amount {
			if data.Amount > amountnhistory {
				balance = balance - (data.Amount - amountnhistory)
			} else {
				balance = balance + (amountnhistory - data.Amount)
			}
		} else {
			fmt.Println("Saldo tidak mencukupi")
		}
	case "debet":
		if data.Amount > amountnhistory {
			balance = balance + (data.Amount - amountnhistory)
		} else {
			balance = balance - (amountnhistory - data.Amount)
		}
	}

	data.BalanceHistory = balance
	accountUpdate.Balance = balance
	balancehistory = balance
	err = s.cashflowRepository.RepoUpdateBalance(ctx, &accountUpdate, paramsIdAccount)
	if err != nil {
		return nil, err
	}
	err = s.cashflowRepository.RepoEditCashFlow(ctx, &data, params, balancehistory)
	if err != nil {
		return nil, err
	}

	fmt.Println(balance)

	return &data, nil
}

func NewCashflowService(repository interfaces.CashflowRepositoryInterface) serviceInterface.CashflowServiceInterface {
	return &Service{
		cashflowRepository: repository,
	}
}
