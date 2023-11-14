package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/products", helloProduct)
	_ = app.Listen(":50051")
}

func helloProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello From Service Product",
	})
}
