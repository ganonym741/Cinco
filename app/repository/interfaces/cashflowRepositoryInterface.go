package interfaces

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowRepositoryInterface interface {
	PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error
	FindByAccount(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error)
	DeleteCashflow(ctx *fiber.Ctx, deleteCashFlow *model.Cashflow, params string) error
	RepoEditCashFlow(ctx *fiber.Ctx, editcashflow *model.Cashflow, params string, balancehistory int) error
	RepoUpdateBalance(ctx *fiber.Ctx, req *model.Account, paramsIdAccount string) error
	GetHistoryandAmountBefore(ctx *fiber.Ctx, params string) (int, string, int, error)
}
