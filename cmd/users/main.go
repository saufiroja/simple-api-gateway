package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/users", helloUser)

	_ = app.Listen(":50052")
}

func helloUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello From Service User",
	})
}
