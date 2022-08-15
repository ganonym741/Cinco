package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	utilities "gitlab.com/cinco/utils"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) FindTotal(userUUID string, startDate time.Time, endDate time.Time) (model.Total, error) {
	var query = "SELECT SUM(CASE WHEN c.type = 'credit' THEN c.amount ELSE 0 END) as credit, " +
		"SUM(CASE WHEN c.type = 'debet' THEN c.amount ELSE 0 END) as debet " +
		"FROM cashflows c " +
		"INNER JOIN accounts a ON c.account_id  = a.id INNER JOIN users u ON a.user_id = u.id " +
		"WHERE u.id = '" + userUUID + "'"

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

func (r Repository) Create(account model.Account) error {
	err := r.Db.Create(&account).Error
	return err
}

func (r Repository) GetBalance(ctx *fiber.Ctx, params string) (int, error) {
	var balance int
	err := r.Db.Raw("SELECT balance FROM public.accounts WHERE id = ?", params).Scan(&balance).Error
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r Repository) UpdateBalance(ctx *fiber.Ctx, params string, balance int) error {
	var account model.Account
	err := r.Db.Model(&account).Where("id = ?", params).Update("balance", balance).Error
	if err != nil {
		return err
	}
	return nil
}

func NewAccountRepository(db *gorm.DB) interfaces.AccountRepositoryInterface {
	return &Repository{
		Db: db,
	}
}
