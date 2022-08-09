package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/service"
	"gitlab.com/cinco/pkg/postgres"
	"gitlab.com/cinco/routes"
)

func main() {
	app := fiber.New()

	postgres.ConnectDB()

	appService := service.NewService()

	routes.AllRouter(app, appService)
	app.Listen(":8000")
}
