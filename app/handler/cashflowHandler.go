package handler

import (
	"strconv"

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
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Periksa kembali inputan anda", "data": nil})
	}

	if body.Type == "debet" || body.Type == "credit" {
		err := h.cashflowService.AddTransaction(ctx, body)
		if err != nil {
			return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Server sedang bermasalah, silahkan coba beberapa saat lagi", "data": nil})
		}
	} else {
		return ctx.Status(501).JSON(fiber.Map{"status": "Failed", "message": "Tipe transasksi salah", "data": nil})
	}

	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "Transaksi baru telah ditambahkan", "data": body})
}

func (h Handler) CashflowHistory(ctx *fiber.Ctx) error {
	startDate, _ := strconv.ParseInt(ctx.Query("startdate"), 10, 64)
	endDate, _ := strconv.ParseInt(ctx.Query("enddate"), 10, 64)
	uuid := ctx.Query("uuid")
	tipe := ctx.Query("type")

	if len(uuid) <= 0 || startDate <= 0 || endDate <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "bad request", "message": "bad request", "data": []string{}})
	}

	cashflows := h.cashflowService.FindTransactionLog(uuid, tipe, startDate, endDate)

	if len(cashflows) <= 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "data not found", "message": "no record found", "data": []string{}})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success", "data": cashflows})
}

func (h Handler) CashflowEdit(ctx *fiber.Ctx) error {
	params := ctx.Params("cashflowId")
	paramsIdAccount := ctx.Params("accountId")
	// paramDate := ctx.Params("date")

	var modelcashflow model.Cashflow
	ctx.BodyParser(&modelcashflow)

	//
	var modelaccount model.Account
	ctx.BodyParser(&modelaccount)

	data, err := h.cashflowService.EditCashflow(ctx, &modelcashflow, &modelaccount, params, paramsIdAccount)
	if err != nil {
		return ctx.Status(200).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}
	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}

func (h Handler) CashflowDelete(ctx *fiber.Ctx) error {
	params := ctx.Params("cashflowId")
	paramsIdAccount := ctx.Params("accountId")

	data, err := h.cashflowService.DeleteCashflow(ctx, params, paramsIdAccount)
	if err != nil {
		return ctx.Status(200).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}
	return ctx.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}

func NewCashflowHandler(service interfaces.CashflowServiceInterface) CincoCashflow {
	return &Handler{
		cashflowService: service,
	}
}
