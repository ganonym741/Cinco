package service

import "gitlab.com/cinco/app/repository"

type UserService struct {
	cashflowRepository repository.CashflowRepository
}
