package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/config"
)

type Helicopter struct {
	Name string `query:"name"`
	Year int32  `query:"year"`
	Type string
}

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
	app.Get("/chopper", getHelicopter)

	port := fmt.Sprintf(":%s", c.Port)

	app.Listen(port)
}

func getName(c *fiber.Ctx) error {
	now := time.Now()

	name := c.Params("name")

	result := fmt.Sprintf("Hello, %+v", name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code":   fiber.StatusOK,
		"message":       result,
		"response_time": time.Since(now).String(),
	})
}

func getHelicopter(c *fiber.Ctx) error {
	hind := Helicopter{
		Name: "Mi-24 Hind",
		Year: 1964,
		Type: "Attack Helicopter",
	}

	apache := Helicopter{
		Name: "AH-64D Apache",
		Year: 1980,
		Type: "Attack Helicopter",
	}

	helicopterName := c.Query("name")

	helicopters := []Helicopter{hind, apache}

	for _, helicopter := range helicopters {
		lowercase := strings.ToLower(helicopter.Name)
		match, _ := regexp.MatchString(helicopterName, lowercase)

		if match {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "success",
				"result":  helicopter,
			})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Helicopter Not Found!",
	})
}
