package repository

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gorm.io/gorm"
)

type CashflowRepository struct {
	Db *gorm.DB
}

func (c CashflowRepository) FindByAccount(userUUID string, startDate int, endDate int) []model.Cashflow {
	query := "SELECT c.type, c.ammount, c.description, c.created_at FROM Cashflow c " +
		"INNER JOIN Account a ON c.account_uuid = a.account_uuid " +
		"INNER JOIN User u ON a.user_uuid = u.user_uuid "
	if len(userUUID) != 0 && userUUID != "" {
		query += "WHERE u.user_uuid = ? " + userUUID
	}

	if startDate > 0 && endDate > 0 {
		query += "AND c.created_at BETWEEN ++ " +
			"AND "
	}
	var cashflows []model.Cashflow
	c.Db.Raw(query).Scan(&cashflows)

	return cashflows
}

func NewCashflowRepository() interfaces.CashflowRepositoryInterface {
	return &CashflowRepository{}
}

/*func (r Repository) PostTransaction(ctx fiber.Ctx, cashflow model.Cashflow) error {
	err := r.Db.Create(&cashflow).Error
	return err
}*/
