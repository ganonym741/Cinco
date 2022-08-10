package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/app/model"
)

func (s Service) AddTransaction(ctx *fiber.Ctx, body *model.Cashflow) error {
	body.Id = uuid.NewString()

	balance, err := s.repository.GetBalance(ctx, body.AccountId)
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

	err = s.repository.UpdateBalance(ctx, body.AccountId, balance)
	if err != nil {
		return err
	}

	err = s.repository.PostTransaction(ctx, body)
	if err != nil {
		return err
	}
	return nil
}
