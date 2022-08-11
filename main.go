package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"gitlab.com/cinco/pkg/postgres"
	"gitlab.com/cinco/routes"
	"io"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db := postgres.ConnectDB()

	routes.AllRouter(app, db)
	app.Listen(":8000")
}

type Views interface {
	Load() error
	Render(io.Writer, string, interface{}, ...string) error
}
