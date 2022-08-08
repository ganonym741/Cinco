package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/pkg/postgres"
	"gitlab.com/cinco/routes"
)

func main() {
	app := fiber.New()

	postgres.ConnectDB()

	routes.AllRouter(app)
	app.Listen(":8000")
}
