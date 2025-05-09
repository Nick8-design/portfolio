package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Nick Dieda",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", nil)
	})

	app.Get("/skills", func(c *fiber.Ctx) error {
		return c.Render("skills", nil)
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.Render("projects", nil)
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", nil)
	})

	app.Get("/app-ads.txt", func(c *fiber.Ctx) error {
		return c.SendString("google.com, pub-5272550552627150, DIRECT, f08c47fec0942fa0")
	})

	log.Fatal(app.Listen(":21888"))
}
