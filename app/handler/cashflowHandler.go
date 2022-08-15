package handler

import (
	"time"

	utilities "gitlab.com/cinco/utils"

	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/service/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	cashflowService interfaces.CashflowServiceInterface
}

type CincoCashflow interface {
	DoTransaction(ctx *fiber.Ctx) error
	CashflowEdit(c *fiber.Ctx) error
	CashflowDelete(c *fiber.Ctx) error
	CashflowHistory(c *fiber.Ctx) error
}

func (h Handler) DoTransaction(ctx *fiber.Ctx) error {
	var body model.Cashflow

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
	}

	if body.Type == "debet" || body.Type == "credit" {
		err := h.cashflowService.AddTransaction(ctx, body)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
		}
	} else {
		return ctx.Status(400).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
	}

	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan"})
}

func (h Handler) CashflowHistory(ctx *fiber.Ctx) error {
	startDate, err := time.Parse(utilities.LayoutFormat, ctx.Query("startdate"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "bad request", "message": "invalid date format", "data": []string{}})
	}
	endDate, err := time.Parse(utilities.LayoutFormat, ctx.Query("enddate"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "bad request", "message": "invalid date format", "data": []string{}})
	}
	uuid := ctx.Query("uuid")
	tipe := ctx.Query("type")

	if len(uuid) <= 0 || startDate.IsZero() || endDate.IsZero() {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "bad request", "message": "bad request data", "data": []string{}})
	}

	cashflows, err := h.cashflowService.FindTransactionLog(uuid, tipe, startDate, endDate)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "internal server error", "message": "error", "data": []string{}})
	}

	if len(cashflows) <= 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "no record found", "data": []string{}})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success", "data": cashflows})
}

func (h Handler) CashflowEdit(ctx *fiber.Ctx) error {
	params := ctx.Params("cashflowId")
	paramsIdAccount := ctx.Params("accountId")

	var modelcashflow model.Cashflow
	ctx.BodyParser(&modelcashflow)

	//
	var modelaccount model.Account
	ctx.BodyParser(&modelaccount)

	data, err := h.cashflowService.EditCashflow(ctx, &modelcashflow, &modelaccount, params, paramsIdAccount)
	if err != nil {
		return ctx.Status(404).
			JSON(fiber.Map{"status": "Failed", "message": "Data not found", "data": nil})
	}
	return ctx.Status(201).
		JSON(fiber.Map{"status": "Success", "message": "User data retrieved", "data": data})
}

func (h Handler) CashflowDelete(ctx *fiber.Ctx) error {
	params := ctx.Params("cashflowId")
	paramsIdAccount := ctx.Params("accountId")

	err := h.cashflowService.DeleteCashflow(ctx, params, paramsIdAccount)
	if err != nil {
		return ctx.Status(404).
			JSON(fiber.Map{"status": "Failed", "message": "Data not found", "data": nil})
	}
	return ctx.Status(201).
		JSON(fiber.Map{"status": "Success", "message": "Cashflow deleted"})
}

func NewCashflowHandler(service interfaces.CashflowServiceInterface) CincoCashflow {
	return &Handler{
		cashflowService: service,
	}
}
