
package router

import (
	"cherry-blog/db/dao/helloworld"
	"cherry-blog/pkg/helloworld/controller"
	"cherry-blog/pkg/helloworld/service"

	"github.com/gofiber/fiber/v2"


	"gorm.io/gorm"
)

func InitializeRouters(app fiber.Router, connection *gorm.DB) {

	dataAccess := helloworld.NewDataAccess(connection)
	service := service.NewService(dataAccess)
	controller := controller.NewController(service) // Assign controller to a variable, so it can be used in validator and router

	// Assign to a variable, use methods like .Get(), .Post() etc.
	api := app.Group("/helloworld")

	// Define routes
	api.Get("/", controller.HelloWorld)

}

