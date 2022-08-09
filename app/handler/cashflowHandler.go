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
func CashflowEdit(c *fiber.Ctx) {

}
func CashflowDelete(c *fiber.Ctx) {

}
func CashflowHistory(c *fiber.Ctx) {

}
