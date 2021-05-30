package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	id int `json:"id"`
	name string `json:"string"`
}

var todo = []Todo{
	{id :1, name: "Dimas"},
	{id :2 ,name: "Dimas"},
}

func main ()  {
	app := fiber.New()

	// app.Use(middleware.Logger())
	
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })


	log.Fatal(app.Listen(":3000"))
}