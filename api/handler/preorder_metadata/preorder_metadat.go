package preorder_metadata

import (
	"fmt"

	c "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	h "github.com/bhanupbalusu/custpreorderms/api/handler_interface"

	"github.com/gofiber/fiber/v2"
)

type preOrderRoutesHandler struct {
	poci c.PreOrderControllerInterface
}

func NewPreOrderRoutesHandler(poci c.PreOrderControllerInterface) h.PreOrderRoutesHandlerInterface {
	return &preOrderRoutesHandler{poci}
}

func (r *preOrderRoutesHandler) GetAll(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	fmt.Println("---------Handler SignUp before calling Controller.SignUp---------")
	return r.poci.GetAll(ctx)
}

func (r *preOrderRoutesHandler) GetOne(ctx *fiber.Ctx) error {
	return r.poci.GetOne(ctx)
}

func (r *preOrderRoutesHandler) Create(ctx *fiber.Ctx) error {
	return r.poci.Create(ctx)
}

func (r *preOrderRoutesHandler) Delete(ctx *fiber.Ctx) error {
	return r.poci.Delete(ctx)
}

func (r *preOrderRoutesHandler) Install(app *fiber.App) {
	app.Get("/preorder", r.GetAll)
	app.Get("/preorder/:id", r.GetOne)
	app.Post("/preorder", r.Create)
	app.Delete("/preorder/:id", r.Delete)
}
