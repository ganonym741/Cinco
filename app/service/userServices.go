package service

import (
	"context"

	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository"
	"gitlab.com/cinco/pkg/postgres"
)

type Service struct {
	repository repository.Repository
}

func NewService() Service {
	return Service{
		repository: repository.Repository{
			Db: postgres.ConnectDB(),
		},
	}
}

func (s Service) GetUserDetail(ctx context.Context, userid string) (*model.User, error) {
	var data *model.User
	return data, nil
}
