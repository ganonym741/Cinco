package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CincoCashflow interface {
	DoTransaction()
	CashflowEdit()
	CashflowDelete()
	CashflowHistory()
}

func (h Handler) DoTransaction(c *fiber.Ctx) error {
	ctx := context.Background()
	var body model.Cashflow

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
	}

	if body.Type == "Uang Masuk" || body.Type == "Uang Keluar" {
		_, err := h.service.AddTransaction(ctx, &body)
		if err != nil {
			return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
		}
	} else {
		return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
	}

	return c.Status(201).
		JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan", "data": cashflow})
}
func CashflowEdit(c *fiber.Ctx) {

}
func CashflowDelete(c *fiber.Ctx) {

}
func CashflowHistory(c *fiber.Ctx) {

}
