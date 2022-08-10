package repository

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gorm.io/gorm"
)

func (r Repository) FindByAccount(userUUID string, tipe string, startDate time.Time, endDate time.Time) []model.Cashflow {
	format := "2006-01-02 15:04:05"

	var query = "SELECT c.id, c.type, c.amount, c.balance_history, c.description " +
		"FROM cashflows c " +
		"INNER JOIN accounts a ON c.account_id  = a.id " +
		"INNER JOIN users u ON a.user_id = u.id " +
		"WHERE u.id = '" + userUUID + "' "

	if len(tipe) > 0 && tipe != "" {
		query += " AND c.type = '" + tipe + "' "
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query += " AND c.created_at BETWEEN '" + Bod(startDate).Format(format) + "' AND '" + Eod(endDate).Format(format) + "'"
	}

	var cashflows []model.Cashflow
	r.Db.Raw(query).Scan(&cashflows)

	return cashflows
}

func (r Repository) PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error {
	err := r.Db.Create(&body).Error
	return err
}

func (r Repository) DeleteCashflow(ctx *fiber.Ctx, deleteCashFlow *model.Cashflow, params string) error {

	err := r.Db.Delete(&deleteCashFlow, "id = ?", params).Error

	return err
}

func (r Repository) RepoEditCashFlow(ctx *fiber.Ctx, editcashflow *model.Cashflow, params string, balancehistory int) error {

	err := r.Db.Model(&model.Cashflow{}).Where("id = ?", params).Updates((map[string]interface{}{"description": editcashflow.Description, "amount": editcashflow.Amount, "balance_history": balancehistory})).Error

	return err
}

func (r Repository) RepoUpdateBalance(ctx *fiber.Ctx, req *model.Account, paramsIdAccount string) error {

	err := r.Db.Model(&model.Account{}).Where("id = ?", paramsIdAccount).Update("balance", req.Balance).Error

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

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

func NewCashflowRepository(db *gorm.DB) interfaces.CashflowRepositoryInterface {
	return &Repository{
		Db: db,
	}
}
