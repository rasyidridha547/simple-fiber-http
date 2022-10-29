package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rasyidridha547/simple-fiber-http/handler"
)

func Routes(app *fiber.App) {
	// root
	app.Get("/", handler.HelloWorld)
	app.Get("/user/:name", handler.Name)

	chopper := app.Group("/chopper", logger.New())

	chopper.Get("/", handler.HelicopterName)
	chopper.Get("/type", handler.HelicopterType)
}
