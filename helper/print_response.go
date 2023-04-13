package helper

import (
	"github.com/MCPutro/golang-todo/model/web"
	"github.com/gofiber/fiber/v2"
)

func PrintResponse(c *fiber.Ctx, statusCode int, status string, message interface{}, data interface{}) error {
	return c.Status(statusCode).JSON(web.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})

}
