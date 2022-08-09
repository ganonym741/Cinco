package service

import (
	"gitlab.com/cinco/app/model"
	repoInterfaces "gitlab.com/cinco/app/repository/interfaces"
	"gitlab.com/cinco/app/service/interfaces"
)

type AccountService struct {
	accountRepository repoInterfaces.AccountRepositoryInterface
}

func (a AccountService) UserActivation(userUUID string) error {
	account := model.Account{
		Balance: 0,
		UserId:  userUUID,
	}

	err := a.accountRepository.Create(account)

	if err == nil {

	}

	return err
}

func NewAccountService(accoutRepository repoInterfaces.AccountRepositoryInterface) interfaces.AccountServiceInterface {
	return &AccountService{
		accountRepository: accoutRepository,
	}
}
