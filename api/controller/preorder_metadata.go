package controller

import (
	"fmt"
	"log"
	"net/http"

	i "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/preorder_metadata"
	u "github.com/bhanupbalusu/custpreorderms/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type PreOrderController struct {
	posi s.PreOrderServiceInterface
}

func NewPreOrderController(posi s.PreOrderServiceInterface) i.PreOrderControllerInterface {
	return &PreOrderController{posi}
}

func (pdc *PreOrderController) GetAll(ctx *fiber.Ctx) error {
	pd, err := pdc.posi.Get()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (pdc *PreOrderController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	fmt.Println(id)
	pd, err := pdc.posi.GetByID(id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (pdc *PreOrderController) Create(ctx *fiber.Ctx) error {
	var req m.PreOrderMetaDataModel
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	fmt.Println("------- Inside Handler Create Method Before Calling ProductService.Create -----------")
	fmt.Println(req)
	pid, err := pdc.posi.Create(&req)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pid)

}

func (pdc *PreOrderController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := pdc.posi.Delete(id); err != nil {
		log.Fatal(err)
	}
	return ctx.SendString("Product Is Deleted")
}
