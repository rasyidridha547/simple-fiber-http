package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/config"
)

func main() {
	c := config.Get()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		now := time.Now()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":       true,
			"message":       "Hello World!",
			"response_time": time.Since(now).String(),
		})
	})

	app.Get("/user/:name", getName)

	app.Listen(fmt.Sprintf(":%s", c.Port))
}

func getName(c *fiber.Ctx) error {
	name := c.Params("name")

	result := fmt.Sprintf("Hello, %+v", name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"message":     result,
	})
}
