package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("./public")

	app.Use(func(c *fiber.Ctx) {
		c.Write("Match anything!\n")
		c.Next()
	})

	app.Use("/api", func(c *fiber.Ctx) {
		c.Write("Match starting with /api\n")
		c.Next()
	})

	app.Get("/api/user", func(c *fiber.Ctx) {
		c.Write("Match exact path /api/user\n")
	})

	app.Listen(7000)
}
