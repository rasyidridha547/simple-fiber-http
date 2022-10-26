package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/config"
)

func main() {
	c := config.Get()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Hello World!",
		})
	})

	app.Listen(fmt.Sprintf(":%s", c.Port))
}
