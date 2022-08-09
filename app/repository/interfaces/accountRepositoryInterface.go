package interfaces

import (
	"gitlab.com/cinco/app/model"
)

type AccountRepositoryInterface interface {
	Create(account model.Account) error
}
