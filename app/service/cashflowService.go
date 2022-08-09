package service

import (
	"context"

	"gitlab.com/cinco/app/model"
)

func (s Service) AddTransaction(ctx context.Context, userid string) (*model.User, error) {
	var data model.User
	err := s.repository.GetUserDetail(ctx, &data, userid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
