package order_types

import (
	"fmt"

	c "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	h "github.com/bhanupbalusu/custpreorderms/api/handler_interface"

	"github.com/gofiber/fiber/v2"
)

type orderTypesRoutesHandler struct {
	otci c.OrderTypesControllerInterface
}

func NewOrderTypesRoutesHandler(otci c.OrderTypesControllerInterface) h.OrderTypesRoutesHandlerInterface {
	return &orderTypesRoutesHandler{otci}
}

func (r *orderTypesRoutesHandler) GetAll(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	fmt.Println("---------Handler SignUp before calling Controller.SignUp---------")
	return r.otci.GetAll(ctx)
}

func (r *orderTypesRoutesHandler) GetOne(ctx *fiber.Ctx) error {
	return r.otci.GetOne(ctx)
}

func (r *orderTypesRoutesHandler) Create(ctx *fiber.Ctx) error {
	return r.otci.Create(ctx)
}

func (r *orderTypesRoutesHandler) Update(ctx *fiber.Ctx) error {
	return r.otci.Update(ctx)
}

func (r *orderTypesRoutesHandler) Delete(ctx *fiber.Ctx) error {
	return r.otci.Delete(ctx)
}

func (r *orderTypesRoutesHandler) Install(app *fiber.App) {
	app.Get("/ordertypes", r.GetAll)
	app.Get("/ordertypes/:id", r.GetOne)
	app.Post("/ordertypes", r.Create)
	app.Put("/ordertypes/:id", r.Update)
	app.Delete("/ordertypes/:id", r.Delete)
}
