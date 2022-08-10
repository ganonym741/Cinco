package handler

import "github.com/gofiber/fiber/v2"

type CincoAccount interface {
	AccountActivation()
}

func (h Handler) AccountActivation(c *fiber.Ctx) error {
	return nil
}
