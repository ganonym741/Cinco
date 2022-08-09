package interfaces

import "gitlab.com/cinco/app/model"

type AccountRepositoryInterface interface {
	Activation(account model.Account) error
}
