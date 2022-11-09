package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/internal"
)

func notFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Query Not Found!",
	})
}

func HelloWorld(c *fiber.Ctx) error {
	now := time.Now()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":       true,
		"message":       "Hello World!",
		"response_time": time.Since(now).String(),
	})
}

func Name(c *fiber.Ctx) error {
	now := time.Now()

	name := c.Params("name")

	notFoundError := notFound(c)

	if name == "" {
		return notFoundError
	}

	result := internal.GetName(name)

	if result == "" {
		return notFoundError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code":   fiber.StatusOK,
		"message":       result,
		"response_time": time.Since(now).String(),
	})
}

func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "service is alive!",
	})
}

func HelicopterType(c *fiber.Ctx) error {
	chopperType := c.Query("name")

	notFoundError := notFound(c)

	if chopperType == "" {
		return notFoundError
	}

	result := internal.TypeHelicopter(chopperType)

	if result == nil {
		return notFoundError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"result":  result,
	})
}

func HelicopterName(c *fiber.Ctx) error {
	chopperName := c.Query("name")

	notFoundError := notFound(c)

	if chopperName == "" {
		return notFoundError
	}

	result := internal.NameHelicopter(chopperName)

	if result == nil {
		return notFoundError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"result":  result,
	})
}
