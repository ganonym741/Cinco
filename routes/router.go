package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/cinco/app/repository"
	"gorm.io/gorm"

	"gitlab.com/cinco/app/handler"
	"gitlab.com/cinco/app/service"
)

func AllRouter(app *fiber.App, db *gorm.DB) {
	cashflowRepository := repository.NewCashflowRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	userRepository := repository.NewUserRepository(db)

	cashflowService := service.NewCashflowService(cashflowRepository)
	accountService := service.NewAccountService(accountRepository)
	userService := service.NewUserService(userRepository)

	cashflowHandler := handler.NewCashflowHandler(cashflowService)
	accountHandler := handler.NewAccountHandler(accountService, userService)
	//userHandler := handler.NewUserHandler(userService)

	//Handler := handler.NewHandler(services)
	api := app.Group("/api", logger.New())

	// api.Post("/user/register", Handler.UserRegister)
	// api.Post("/user/login", Handler.UserLogin)
	// api.Post("/user/logout", Handler.UserLogout)
	//api.Get("/user/profile/:userId", Handler.UserProfile)

	api.Post("/user/activation/:userId", accountHandler.AccountActivation)

	//api.Post("/cash", Handler.DoTransaction)
	//api.Put("/user/:cashflowId", Handler.CashflowEdit)
	//api.Delete("/user/:cashflowId", Handler.CashflowDelete)
	api.Get("/cash", cashflowHandler.CashflowHistory)
}
