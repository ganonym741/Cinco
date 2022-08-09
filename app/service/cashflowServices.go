package service

import (
	"gitlab.com/cinco/app/repository/interfaces"
)

type CashflowService struct {
	cashflowRepository interfaces.CashflowRepositoryInterface
}
