package initialize

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	return app
}
