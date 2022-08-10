package interfaces

import (
	"gitlab.com/cinco/app/model"
)

type UserRepositoryInterface interface {
	FindById(userUUID string) model.User
	Update(user model.User) error
}
