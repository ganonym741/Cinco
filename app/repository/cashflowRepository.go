package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

func (r Repository) PostTransaction(ctx fiber.Ctx, cashflow model.Transaction_log) error {
	err := r.Db.Create(&cashflow).Error
	return err
}
