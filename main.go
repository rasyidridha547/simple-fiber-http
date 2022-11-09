package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rasyidridha547/simple-fiber-http/config"
	"github.com/rasyidridha547/simple-fiber-http/router"
)

func main() {
	c := config.Get()

	app := fiber.New()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		_ = <-ch
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	router.Routes(app)

	port := fmt.Sprintf(":%s", c.Port)
	if err := app.Listen(port); err != nil {
		log.Panic(err)
	}
}
