package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gorm.io/gorm"
)

type AccoutRepository struct {
	Db *gorm.DB
}

func (a AccoutRepository) Create(ctx fiber.Ctx, account model.Account) error {
	err := a.Db.Create(&account).Error
	return err
}

func NewAccountRepository() interfaces.AccountRepositoryInterface {
	return &AccoutRepository{}
}
