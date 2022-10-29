package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/config"
	"github.com/rasyidridha547/simple-fiber-http/router"
)

func main() {
	c := config.Get()

	app := fiber.New()
	router.Routes(app)

	port := fmt.Sprintf(":%s", c.Port)
	app.Listen(port)
}
