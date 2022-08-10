package repository

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (u UserRepository) Update(user model.User) error {
	return u.Db.Model(user).Save(user).Error
}

func (u UserRepository) FindById(userUUID string) model.User {
	var user model.User

	u.Db.Where("id = ?", userUUID).First(&user)

	return user
}

/*func (r UserRepository) GetUserDetail(ctx context.Context, user *model.User, params string) error {
	err := r.Db.First(&user, "id = ?", params).Error
	return err
}*/

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{
		Db: db,
	}
}
