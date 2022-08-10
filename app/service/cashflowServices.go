package service

import (
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

func NewCashflowService(repository interfaces.CashflowRepositoryInterface) serviceInterface.CashflowServiceInterface {
	return &Service{
		cashflowRepository: repository,
	}
}
