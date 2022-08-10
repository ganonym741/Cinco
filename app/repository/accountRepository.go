package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

func (r Repository) Create(ctx *fiber.Ctx, account model.Account) error {
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
