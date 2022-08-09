package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"gitlab.com/cinco/app/handler"
	"gitlab.com/cinco/app/service"
)

func AllRouter(app *fiber.App, service service.Service) {
	Handler := handler.NewHandler(service)
	api := app.Group("/api", logger.New())

	// api.Post("/user/register", Handler.UserRegister)
	// api.Post("/user/login", Handler.UserLogin)
	// api.Post("/user/logout", Handler.UserLogout)
	api.Get("/user/profile/:userId", Handler.UserProfile)

	// api.Post("/user/activation", Handler.AccountActivation)

	// api.Post("/cash", Handler.DoTransaction)
	// api.Put("/user/:cashflowId", Handler.CashflowEdit)
	// api.Delete("/user/:cashflowId", Handler.CashflowDelete)
	// api.Get("/cash", Handler.CashflowHistory)
}
