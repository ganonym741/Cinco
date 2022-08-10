package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

func (r Repository) PostTransaction(ctx fiber.Ctx, cashflow model.Cashflow) error {
	err := r.Db.Create(&cashflow).Error
	return err
}

func (r Repository) DeleteCashflow(ctx context.Context, deleteCashFlow *model.Cashflow, params string) error {

	err := r.Db.Delete(&deleteCashFlow, "id = ?", params).Error

	return err
}

func (r Repository) RepoEditCashFlow(ctx context.Context, editcashflow *model.Cashflow, params string) error {

	err := r.Db.Model(&model.Cashflow{}).Where("id = ?", params).Updates((map[string]interface{}{"description": editcashflow.Description, "amount": editcashflow.Amount, "balance_history": editcashflow.BalanceHistory})).Error

	return err

}

func (r Repository) GetBalance(ctx context.Context, params string) (int, error) {

	var result int
	err := r.Db.Table("accounts").Select("balance").Where("user_id = ?", params).Scan(&result).Error

	fmt.Println("ini result", result)
	return result, err
}

func (r Repository) RepoUpdateBalance(ctx context.Context, req *model.Account, params string) error {

	err := r.Db.Model(&model.Account{}).Where("user_id = ?", params).Update("balance", req.Balance).Error

	return err
}

func (r Repository) GetHistoryandAmountBefore(ctx context.Context, params string) (int, string, error) {
	var Result struct {
		Amount int    `json:"amount"`
		Type   string `json:"type"`
	}

	// err := r.Db.Table("cashflows").Select("amount", "type").Where("user_id = ?", params).Scan(result).Error
	err := r.Db.Raw("SELECT amount, type FROM cashflows WHERE account_id = ?", params).Scan(&Result).Error
	fmt.Println("err history before", err)
	fmt.Println("result2", Result)

	return Result.Amount, Result.Type, err
}
