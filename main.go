package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/pkg/postgres"
	"gitlab.com/cinco/routes"
)

func main() {
	app := fiber.New()

	db := postgres.ConnectDB()

	routes.AllRouter(app, db)
	app.Listen(":8000")
}
