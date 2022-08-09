package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) GetUserDetail(ctx fiber.Ctx, user *model.User) error {
	err := r.Db.Find(&user).Error
	return err
}
