package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/app/model"
	repoInterfaces "gitlab.com/cinco/app/repository/interfaces"
	"gitlab.com/cinco/app/service/interfaces"
)

type AccountService struct {
	accountRepository repoInterfaces.AccountRepositoryInterface
}

func (a AccountService) CreateAccount(userUUID string) error {
	account := model.Account{
		Id:      uuid.New().String(),
		Balance: 0,
		UserId:  userUUID,
	}

	return a.accountRepository.Create(account)
}

func (a AccountService) GetBalance(ctx *fiber.Ctx, params string) (int, error) {
	balance, err := a.accountRepository.GetBalance(ctx, params)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func NewAccountService(accoutRepository repoInterfaces.AccountRepositoryInterface) interfaces.AccountServiceInterface {
	return &AccountService{
		accountRepository: accoutRepository,
	}
}
