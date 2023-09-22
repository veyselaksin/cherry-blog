package controller

import (
	"cherry-blog/pkg/helloworld/service"

	"cherry-blog/utils/response"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	// Define controller methods
	HelloWorld(c *fiber.Ctx) error
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return controller{service: service}
}

func (c controller) HelloWorld(ctx *fiber.Ctx) error {

	data, statusCode, err := c.service.Helloworld()
	if err != nil {
		return response.ErrorResponse(ctx, statusCode, err.Error())
	}

	return response.SuccessResponse(ctx, data, 200)
}
