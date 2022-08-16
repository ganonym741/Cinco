package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/cinco/app/repository"
	"gorm.io/gorm"

	"gitlab.com/cinco/app/handler"
	"gitlab.com/cinco/app/service"
	utilities "gitlab.com/cinco/utils"
)

func AllRouter(app *fiber.App, db *gorm.DB) {
	cashflowRepository := repository.NewCashflowRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	userRepository := repository.NewUserRepository(db)

	cashflowService := service.NewCashflowService(cashflowRepository, accountRepository)
	accountService := service.NewAccountService(accountRepository)
	userService := service.NewUserService(userRepository)

	cashflowHandler := handler.NewCashflowHandler(cashflowService)
	accountHandler := handler.NewAccountHandler(accountService, userService)
	userHandler := handler.NewUserHandler(userService)

	//Handler := handler.NewHandler(services)
	api := app.Group("/api", logger.New())

	api.Post("/user/register", userHandler.UserRegister)
	api.Post("/user/login", userHandler.UserLogin)
	api.Get("/user/activation/:userId", accountHandler.AccountActivation)

	app.Use(utilities.TokenVerify())

	api.Get("/account/balance", utilities.Authorization(true), accountHandler.GetBalance)

	api.Get("/user/profile", utilities.Authorization(true), userHandler.UserProfile)
	api.Post("/user/logout", utilities.Authorization(true), userHandler.UserLogout)

	api.Get("/cash", utilities.Authorization(true), cashflowHandler.CashflowHistory)
	api.Post("/cash", utilities.Authorization(true), cashflowHandler.DoTransaction)
	api.Put("/cash/:cashflowId/:accountId", utilities.Authorization(true), cashflowHandler.CashflowEdit)
	api.Delete("/cash/:cashflowId/:accountId", utilities.Authorization(true), cashflowHandler.CashflowDelete)
}
