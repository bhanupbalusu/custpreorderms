package product_details

import (
	"fmt"

	c "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	h "github.com/bhanupbalusu/custpreorderms/api/handler_interface"

	"github.com/gofiber/fiber/v2"
)

type dealSettingsRoutesHandler struct {
	dsci c.DealSettingsControllerInterface
}

func NewDealSettingsRoutesHandler(dsci c.DealSettingsControllerInterface) h.DealSettingsRoutesHandlerInterface {
	return &dealSettingsRoutesHandler{dsci}
}

func (r *dealSettingsRoutesHandler) GetAll(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	fmt.Println("---------Handler SignUp before calling Controller.SignUp---------")
	return r.dsci.GetAll(ctx)
}

func (r *dealSettingsRoutesHandler) GetOne(ctx *fiber.Ctx) error {
	return r.dsci.GetOne(ctx)
}

func (r *dealSettingsRoutesHandler) Create(ctx *fiber.Ctx) error {
	return r.dsci.Create(ctx)
}

func (r *dealSettingsRoutesHandler) Update(ctx *fiber.Ctx) error {
	return r.dsci.Update(ctx)
}

func (r *dealSettingsRoutesHandler) Delete(ctx *fiber.Ctx) error {
	return r.dsci.Delete(ctx)
}

func (r *dealSettingsRoutesHandler) Install(app *fiber.App) {
	app.Get("/dealsettings", r.GetAll)
	app.Get("/dealsettings/:id", r.GetOne)
	app.Post("/dealsettings", r.Create)
	app.Put("/dealsettings", r.Update)
	app.Delete("/dealsettings/:id", r.Delete)
}
