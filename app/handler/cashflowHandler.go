package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CincoCashflow interface {
	DoTransaction()
	CashflowEdit()
	CashflowDelete()
	CashflowHistory()
}

func (h Handler) DoTransaction(ctx *fiber.Ctx) error {
	var body model.Cashflow

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
	}

	if body.Type == "debet" || body.Type == "kredit" {
		err := h.service.AddTransaction(ctx, &body)
		if err != nil {
			return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
		}
	} else {
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
	}

	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan", "data": body})
}
func CashflowEdit(c *fiber.Ctx) {

}
func CashflowDelete(c *fiber.Ctx) {

}
func CashflowHistory(c *fiber.Ctx) {

}
