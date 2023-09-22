
package config

import (
	"cherry-blog/utils/response"

	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	AppName:   "cherry-blog API",
	BodyLimit: 1024 * 1024 * 50, // 50 MB

	// Override default error controller
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		var code int = fiber.StatusInternalServerError

		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// log.Fatal("Error occurred: ", goErrors.Wrap(err, 0).ErrorStack())

		return response.ErrorResponse(ctx, code, "Unexpected error occurred")
	},
}
