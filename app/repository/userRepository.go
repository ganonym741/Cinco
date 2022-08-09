package repository

import (
	"context"

	"gitlab.com/cinco/app/model"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) GetUserDetail(ctx context.Context, user *model.User) error {
	err := r.Db.Find(&user).Error
	return err
}
