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

// func DoTransaction(c *fiber.Ctx) error {
// 	db := postgres.DB
// 	var cashflow model.Cashflow
// 	cashflow.CashflowId = uuid.NewString()

// 	err := c.BodyParser(&cashflow)
// 	if err != nil {
// 		return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
// 	}

// 	if cashflow.Type == "Uang Masuk" || cashflow.Type == "Uang Keluar" {
// 		result := db.Create(&cashflow)
// 		if result.Error != nil {
// 			return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
// 		}
// 	} else {
// 		return c.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
// 	}

// 	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan", "data": cashflow})
// }
func (h Handler) CashflowEdit(c *fiber.Ctx) error {
	ctx := context.Background()
	params := c.Params("cashflowId")
	paramsIdAccount := c.Params("accountId")

	var modelcashflow model.Cashflow
	c.BodyParser(&modelcashflow)

	//
	var modelaccount model.Account
	c.BodyParser(&modelaccount)

	data, err := h.service.EditCashflow(ctx, &modelcashflow,&modelaccount, params, paramsIdAccount)
	if err != nil {
		return c.Status(200).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}
	return c.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}

func (h Handler) CashflowDelete(c *fiber.Ctx) error {
	ctx := context.Background()
	params := c.Params("cashflowId")

	data, err := h.service.DeleteCashflow(ctx, params)
	if err != nil {
		return c.Status(200).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}
	return c.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}


