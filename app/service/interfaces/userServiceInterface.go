package interfaces

import "gitlab.com/cinco/app/model"

type UserServiceInterface interface {
	FindByID(userUUID string) model.User
	Update(user model.User) error
}
