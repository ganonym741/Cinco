package service

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type UserService struct {
	userRepository interfaces.UserRepositoryInterface
}

func (u UserService) Update(user model.User) error {
	return u.userRepository.Update(user)
}

func (u UserService) FindByID(userUUID string) model.User {
	return u.userRepository.FindById(userUUID)
}

func NewUserService(repository interfaces.UserRepositoryInterface) serviceInterface.UserServiceInterface {
	return &UserService{
		userRepository: repository,
	}
}

/*func (s Service) GetUserDetail(ctx context.Context, userid string) (*model.User, error) {
	var data model.User
	err := s.repository.GetUserDetail(ctx, &data, userid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}*/
