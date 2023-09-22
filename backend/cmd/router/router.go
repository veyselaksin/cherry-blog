package router

import (
	"cherry-blog/cmd/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeRouters(app fiber.Router, connection *gorm.DB) {

	// Health check routes
	root := app.Group("/")
	controller.HealthCheck(root)

}
