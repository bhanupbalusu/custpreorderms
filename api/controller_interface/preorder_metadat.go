package controller_interface

import (
	"github.com/gofiber/fiber/v2"
)

type PreOrderControllerInterface interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
