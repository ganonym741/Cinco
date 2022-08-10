package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) UserRegister(ctx *fiber.Ctx, params model.User) error {
	err := r.Db.Create(&params).Error
	return err
}

func (r Repository) GetUserByIdentity(ctx *fiber.Ctx, params string) (*model.User, error) {
	var user *model.User
	err := r.Db.Where("username = ? or email = ?", params, params).Find(&user).Error
	return user, err
}
func (r Repository) GetUserDetail(ctx *fiber.Ctx, user *model.User, params string) error {
	err := r.Db.First(&user, "id = ?", params).Error
	return err
}
