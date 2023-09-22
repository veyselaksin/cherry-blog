
package controller

import (
	"cherry-blog/pkg/blog/service"

)

type Controller interface {
	// Define controller methods
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return controller{service: service}
}

