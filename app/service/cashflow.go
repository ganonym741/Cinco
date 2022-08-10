package service

import (
	"context"
	"fmt"

	"gitlab.com/cinco/app/model"
)

func (s Service) DeleteCashflow(ctx context.Context, cashflowid string) (*model.Cashflow, error) {
	var data model.Cashflow

	err := s.repository.DeleteCashflow(ctx, &data, cashflowid)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (s Service) EditCashflow(ctx context.Context, body *model.Cashflow, reqUpdate *model.Account, params string) (*model.Cashflow, error) {

	data := model.Cashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}

	amountnhistory, amounttypes, err := s.repository.GetHistoryandAmountBefore(ctx, params)
	if err != nil {
		return nil, err
	}
	fmt.Println("ini amount history", amountnhistory)

	err = s.repository.RepoEditCashFlow(ctx, &data, params)
	if err != nil {
		return nil, err
	}

	balance, err := s.repository.GetBalance(ctx, params)
	if err != nil {
		return nil, err
	}


	accountUpdate := model.Account{}
	totalAmount := 0

	fmt.Println("ini type", amountnhistory)
	switch amounttypes {
	case "kredit":
		if balance > data.Amount {
			if data.Amount > amountnhistory {
				balance = balance - (data.Amount - amountnhistory)
			} else {
				balance = balance + (amountnhistory - data.Amount)
			}
		} else {
			fmt.Println("saldo tidak mencukupi")
		}
	case "debet":
		if data.Amount > amountnhistory {
			balance = balance + (data.Amount - amountnhistory)
		} else {
			balance = balance - (amountnhistory - data.Amount)
		}
	}

	fmt.Println("ini total amount", totalAmount)
	data.BalanceHistory = balance
	accountUpdate.Balance = balance
	err = s.repository.RepoUpdateBalance(ctx, &accountUpdate, params)

	if err != nil {
		return nil, err
	}
	err = s.repository.RepoEditCashFlow(ctx, &data, params)
	if err != nil {
		return nil, err
	}
	fmt.Println(balance)

	return &data, nil
}
