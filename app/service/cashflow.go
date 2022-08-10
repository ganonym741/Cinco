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

func (s Service) EditCashflow(ctx context.Context, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*model.Cashflow, error) {

	data := model.Cashflow{
		Description: body.Description,
		Amount:      body.Amount,
	}

	fmt.Println("ini data input", data.Amount)

	amountnhistory, amounttypes, balancehistory, err := s.repository.GetHistoryandAmountBefore(ctx, params)
	if err != nil {
		return nil, err
	}

	fmt.Println("S ini amount history", amountnhistory)

	balance, err := s.repository.GetBalance(ctx, paramsIdAccount)
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
			fmt.Println("S saldo tidak mencukupi")
		}
	case "debet":
		if data.Amount > amountnhistory {
			balance = balance + (data.Amount - amountnhistory)
		} else {
			balance = balance - (amountnhistory - data.Amount)
		}
	}

	fmt.Println("S ini total amount")
	data.BalanceHistory = balance
	accountUpdate.Balance = balance
	balancehistory = balance
	err = s.repository.RepoUpdateBalance(ctx, &accountUpdate, paramsIdAccount)
	if err != nil {
		return nil, err
	}
	err = s.repository.RepoEditCashFlow(ctx, &data, params, balancehistory)
	if err != nil {
		return nil, err
	}

	fmt.Println(balance)

	return &data, nil
}
