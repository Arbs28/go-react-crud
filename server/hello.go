package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Print("Hello world")

	app := fiber.New()

	app.Get("/test" ,func(c *fiber.Ctx) error{
		return c.SendString("Ok")
	})

	// log.Fatal(app.Listen(":4000"))
	app.Listen(":4000")
}