package router

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {

	base := app.Group("api")

	base.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

}
