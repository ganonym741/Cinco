package handler

import (
	"gitlab.com/cinco/app/service/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CincoCashflow interface {
	DoTransaction()
	CashflowEdit()
	CashflowDelete()
	CashflowHistory(c *fiber.Ctx) error
}

type CashflowHandler struct {
	cashflowService interfaces.CashflowServiceInterface
}

func (ch CashflowHandler) DoTransaction() {
	/*
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
	*/
	panic("implement me")
}

func (ch CashflowHandler) CashflowEdit() {
	//TODO implement me
	panic("implement me")
}

func (ch CashflowHandler) CashflowDelete() {
	//TODO implement me
	panic("implement me")
}

func (ch CashflowHandler) CashflowHistory(c *fiber.Ctx) error {
	startDate, _ := strconv.ParseInt(c.Query("startdate"), 10, 64)
	endDate, _ := strconv.ParseInt(c.Query("enddate"), 10, 64)
	uuid := c.Query("uuid")
	tipe := c.Query("type")

	if len(uuid) <= 0 || startDate <= 0 || endDate <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "bad request", "data": nil})
	}

	cashflows := ch.cashflowService.FindTransactionLog(uuid, tipe, startDate, endDate)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "success", "data": cashflows})
}

func NewCashflowHandler(service interfaces.CashflowServiceInterface) CincoCashflow {
	return &CashflowHandler{
		cashflowService: service,
	}
}
