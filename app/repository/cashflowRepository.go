package repository

import (
	"fmt"
	"time"

	utilities "gitlab.com/cinco/utils"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gorm.io/gorm"
)

func (r Repository) FindByAccount(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error) {
	var query = "SELECT c.id, c.type, c.amount, c.balance_history, c.description, c.issued_at, c.created_at, c.updated_at " +
		"FROM cashflows c " +
		"INNER JOIN accounts a ON c.account_id  = a.id INNER JOIN users u ON a.user_id = u.id " +
		"WHERE u.id = '" + userUUID + "'"

	if len(tipe) > 0 && tipe != "" {
		query += " AND c.type = '" + tipe + "' "
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query += " AND c.created_at BETWEEN '" + startDate.Format(utilities.DateTimeFormat) + "' AND '" + endDate.Format(utilities.DateTimeFormat) + "'"
	}

	query += " ORDER BY c.issued_at"
	//fmt.Println(query)

	var cashflows []model.Cashflow

	err := r.Db.Raw(query).Scan(&cashflows).Error
	if err != nil {
		return nil, err
	}

	return cashflows, nil
}

func (r Repository) PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error {
	fmt.Println("ini beody", body)
	err := r.Db.Create(body).Error
	fmt.Println(err)
	return err
}

func (r Repository) DeleteCashflow(ctx *fiber.Ctx, cashflowid string) error {

	err := r.Db.Raw("DELETE FROM public.cashflows WHERE id = ?", cashflowid).Error

	return err
}

func (r Repository) RepoEditCashFlow(ctx *fiber.Ctx, editcashflow *model.Cashflow, params string, balancehistory int) error {

	err := r.Db.Model(&model.Cashflow{}).Where("id = ?", params).Updates((map[string]interface{}{"description": editcashflow.Description, "amount": editcashflow.Amount, "balance_history": balancehistory, "updated_at": time.Now()})).Error

	return err
}

func (r Repository) RepoUpdateBalance(ctx *fiber.Ctx, req *model.Account, paramsIdAccount string) error {

	err := r.Db.Model(&model.Account{}).Where("id = ?", paramsIdAccount).Update("balance", req.Balance).Error

	return err
}

func (r Repository) RepoUpdateBalance2(ctx *fiber.Ctx, amountAffected int, paramsIdAccount string) error {

	err := r.Db.Raw("UPDATE public.accounts SET balance = balance + ?", amountAffected).Error

	return err
}

func (r Repository) GetHistoryandAmountBefore(ctx *fiber.Ctx, params string) (int, string, int, error) {
	var Result struct {
		Amount         int    `json:"amount"`
		Type           string `json:"type"`
		BalanceHistory int    `json:"balance_history"`
	}

	// err := r.Db.Table("cashflows").Select("amount", "type").Where("user_id = ?", params).Scan(result).Error
	err := r.Db.Raw("SELECT amount, type, balance_history FROM cashflows WHERE id = ?", params).Scan(&Result).Error
	fmt.Println("err history before", err)
	fmt.Println("result2", Result)

	return Result.Amount, Result.Type, Result.BalanceHistory, nil
}

func NewCashflowRepository(db *gorm.DB) interfaces.CashflowRepositoryInterface {
	return &Repository{
		Db: db,
	}
}
