package handler

import (
	"strconv"

	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/service/interfaces"

	"github.com/gofiber/fiber/v2"
)

type CincoCashflow interface {
	DoTransaction(ctx *fiber.Ctx) error
	CashflowEdit(c *fiber.Ctx) error
	CashflowDelete(c *fiber.Ctx) error
	CashflowHistory(c *fiber.Ctx) error
}

type Handler struct {
	cashflowService interfaces.CashflowServiceInterface
}

func (h Handler) DoTransaction(ctx *fiber.Ctx) error {
	var body model.Cashflow

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
	}

	if body.Type == "debet" || body.Type == "kredit" {
		err := h.cashflowService.AddTransaction(ctx, &body)
		if err != nil {
			return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
		}
	} else {
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
	}

	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan", "data": body})
}

func (h Handler) CashflowEdit(c *fiber.Ctx) error {
	return nil
}

func (h Handler) CashflowDelete(c *fiber.Ctx) error {
	return nil
}

func (h Handler) CashflowHistory(c *fiber.Ctx) error {
	startDate, _ := strconv.ParseInt(c.Query("startdate"), 10, 64)
	endDate, _ := strconv.ParseInt(c.Query("enddate"), 10, 64)
	uuid := c.Query("uuid")
	tipe := c.Query("type")

	if len(uuid) <= 0 || startDate <= 0 || endDate <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "bad request", "data": nil})
	}

	cashflows := h.cashflowService.FindTransactionLog(uuid, tipe, startDate, endDate)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "success", "data": cashflows})
}

func NewCashflowHandler(service interfaces.CashflowServiceInterface) CincoCashflow {
	return &Handler{
		cashflowService: service,
	}
}
