package handler

import (
	"github.com/gofiber/fiber/v2"
)

type CincoCashflow interface {
	DoTransaction()
	CashflowEdit()
	CashflowDelete()
	CashflowHistory()
}

func CashflowEdit(c *fiber.Ctx) {

}
func CashflowDelete(c *fiber.Ctx) {

}
func CashflowHistory(c *fiber.Ctx) {

}
