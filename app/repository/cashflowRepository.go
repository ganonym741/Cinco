package repository

import (
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
