package service

import (
	"gitlab.com/cinco/app/model"
	repoInterfaces "gitlab.com/cinco/app/repository/interfaces"
	"gitlab.com/cinco/app/service/interfaces"
)

type AccountService struct {
	accountRepository repoInterfaces.AccountRepositoryInterface
}

func (a AccountService) CreateAccount(userUUID string) error {
	account := model.Account{
		Balance: 0,
		UserId:  userUUID,
	}

	return a.accountRepository.Create(account)
}

func NewAccountService(accoutRepository repoInterfaces.AccountRepositoryInterface) interfaces.AccountServiceInterface {
	return &AccountService{
		accountRepository: accoutRepository,
	}
}
