package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rasyidridha547/simple-fiber-http/handler"
)

func Routes(app *fiber.App) {
	// root
	root := app.Group("/", logger.New())

	root.Get("", handler.HelloWorld)
	root.Get("/user/:name", handler.Name)
	root.Get("/health", handler.Health)

	chopper := app.Group("/chopper", logger.New())

	chopper.Get("/", handler.HelicopterName)
	chopper.Get("/type", handler.HelicopterType)
}
