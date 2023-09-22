
package controller

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(app fiber.Router) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to cherry-blog API")
	})

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("cherry-blog is running")
	})
}
