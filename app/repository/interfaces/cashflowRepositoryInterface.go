package interfaces

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowRepositoryInterface interface {
	PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error
	FindByAccount(userUUID string, tipe string, startDate time.Time, endDate time.Time) []model.Cashflow
}
