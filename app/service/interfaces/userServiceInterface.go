package interfaces

import "gitlab.com/cinco/app/model"

type UserServiceInterface interface {
	FindByID(userUUID string) model.Cashflow
}
