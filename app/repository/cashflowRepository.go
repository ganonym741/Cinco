package repository

import (
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
		query += " AND c.issued_at BETWEEN '" + startDate.Format(utilities.DateTimeFormat) + "' AND '" + endDate.Format(utilities.DateTimeFormat) + "'"
	}

	query += " AND c.deleted_at IS NULL ORDER BY c.issued_at"

	var cashflows []model.Cashflow

	err := r.Db.Raw(query).Scan(&cashflows).Error
	if err != nil {
		return nil, err
	}

	return cashflows, nil
}

func (r Repository) PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error {
	err := r.Db.Create(body).Error
	return err
}

func (r Repository) DeleteCashflow(ctx *fiber.Ctx, params string) error {
	var cashflow model.Cashflow

	err := r.Db.Where("id = ?", params).Delete(&cashflow).Error

	return err
}

func (r Repository) RepoEditCashFlow(ctx *fiber.Ctx, editcashflow *model.Cashflow, params string) error {

	err := r.Db.Model(&model.Cashflow{}).Where("id = ?", params).Updates((map[string]interface{}{"description": editcashflow.Description, "amount": editcashflow.Amount, "updated_at": time.Now()})).Error

	return err
}

func (r Repository) RepoUpdateBalance(ctx *fiber.Ctx, updatebalance int, paramsIdAccount string) error {

	err := r.Db.Model(&model.Account{}).Where("id = ?", paramsIdAccount).Update("balance", updatebalance).Error

	return err
}

func (r Repository) GetHistoryandAmountBefore(ctx *fiber.Ctx, params string) (int, string, error) {
	var Result struct {
		Amount int    `json:"amount"`
		Type   string `json:"type"`
	}

	err := r.Db.Raw("SELECT amount, type, balance_history FROM cashflows WHERE id = ?", params).Scan(&Result).Error
	if err != nil {
		return 0, "", err
	}

	return Result.Amount, Result.Type, nil
}

func (r Repository) FindTotal(userUUID string, startDate time.Time, endDate time.Time) (model.Total, error) {
	var query = "SELECT SUM(CASE WHEN c.type = 'credit' THEN c.amount ELSE 0 END) as credit, " +
		"SUM(CASE WHEN c.type = 'debet' THEN c.amount ELSE 0 END) as debet " +
		"FROM cashflows c " +
		"INNER JOIN accounts a ON c.account_id  = a.id INNER JOIN users u ON a.user_id = u.id " +
		"WHERE u.id = '" + userUUID + "' AND c.deleted_at IS NULL "

	if !startDate.IsZero() && !endDate.IsZero() {
		query += " AND c.issued_at BETWEEN '" + startDate.Format(utilities.DateTimeFormat) + "' AND '" + endDate.Format(utilities.DateTimeFormat) + "'"
	}

	var totals = model.Total{Debet: 0, Credit: 0}

	err := r.Db.Raw(query).Scan(&totals).Error
	if err != nil {
		return totals, err
	}

	return totals, nil
}

func NewCashflowRepository(db *gorm.DB) interfaces.CashflowRepositoryInterface {
	return &Repository{
		Db: db,
	}
}
