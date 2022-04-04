package handler_interface

import (
	"github.com/gofiber/fiber/v2"
)

type UserAuthRoutesHandlerInterface interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	PutUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	Install(app *fiber.App)
}
