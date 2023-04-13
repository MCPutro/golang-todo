package controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
