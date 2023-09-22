
package router

import (
	"cherry-blog/db/dao/blog"
	"cherry-blog/pkg/blog/controller"
	"cherry-blog/pkg/blog/service"

	"github.com/gofiber/fiber/v2"


	"gorm.io/gorm"
)

func InitializeRouters(app fiber.Router, connection *gorm.DB) {

	dataAccess := blog.NewDataAccess(connection)
	service := service.NewService(dataAccess)
	controller.NewController(service) // Assign controller to a variable, so it can be used in validator and router

	// Assign to a variable, use methods like .Get(), .Post() etc.
	app.Group("/blog") 

}

