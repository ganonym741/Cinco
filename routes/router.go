package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/cinco/app/handler"
	"gitlab.com/cinco/app/service"
	utilities "gitlab.com/cinco/utils"
)

func AllRouter(app *fiber.App, service service.Service) {
	Handler := handler.NewHandler(service)
	api := app.Group("/api", logger.New())

	api.Post("/user/register", Handler.UserRegister)
	api.Post("/user/login", Handler.UserLogin)

	app.Use(utilities.TokenVerify())

	api.Get("/user/profile", utilities.Authorization(true), Handler.UserProfile)
	api.Post("/user/logout", utilities.Authorization(true), Handler.UserLogout)
	api.Post("/cash", utilities.Authorization(true), Handler.DoTransaction)
}
