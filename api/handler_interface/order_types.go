package handler_interface

import (
	"github.com/gofiber/fiber/v2"
)

type OrderTypesRoutesHandlerInterface interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Install(app *fiber.App)
}
