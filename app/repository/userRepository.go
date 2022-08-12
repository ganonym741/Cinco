package repository

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (u UserRepository) UserRegister(ctx *fiber.Ctx, params model.User) error {
	err := u.Db.Create(&params).Error
	return err
}

func (u UserRepository) GetUserByIdentity(ctx *fiber.Ctx, params string) (*model.User, error) {
	var user *model.User
	err := u.Db.Where("username = ? or email = ?", params, params).Find(&user).Error
	return user, err
}

func (u UserRepository) Update(user model.User) error {
	return u.Db.Model(user).Save(user).Error
}

func (u UserRepository) FindById(userUUID string) model.User {
	var user model.User

	u.Db.Where("id = ?", userUUID).First(&user)

	return user
}

func (u UserRepository) GetUserDetail(ctx *fiber.Ctx, user *model.User, params string) error {
	err := u.Db.First(&user, "id = ?", params).Error
	return err
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{
		Db: db,
	}
}
