package product_details

import (
	"fmt"

	c "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	h "github.com/bhanupbalusu/custpreorderms/api/handler_interface"

	"github.com/gofiber/fiber/v2"
)

type productDetailsRoutesHandler struct {
	pdci c.ProductDetailsControllerInterface
}

func NewProductDetailsRoutesHandler(pdci c.ProductDetailsControllerInterface) h.ProductDetailsRoutesHandlerInterface {
	return &productDetailsRoutesHandler{pdci}
}

func (r *productDetailsRoutesHandler) GetAll(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	fmt.Println("---------Handler SignUp before calling Controller.SignUp---------")
	return r.pdci.GetAll(ctx)
}

func (r *productDetailsRoutesHandler) GetOne(ctx *fiber.Ctx) error {
	return r.pdci.GetOne(ctx)
}

func (r *productDetailsRoutesHandler) Create(ctx *fiber.Ctx) error {
	return r.pdci.Create(ctx)
}

func (r *productDetailsRoutesHandler) Update(ctx *fiber.Ctx) error {
	return r.pdci.Update(ctx)
}

func (r *productDetailsRoutesHandler) Delete(ctx *fiber.Ctx) error {
	return r.pdci.Delete(ctx)
}

func (r *productDetailsRoutesHandler) Install(app *fiber.App) {
	app.Get("/productdetails", r.GetAll)
	app.Get("/productdetails/:id", r.GetOne)
	app.Post("/productdetails", r.Create)
	app.Put("/productdetails", r.Update)
	app.Delete("/productdetails/:id", r.Delete)
}
