package controller

import (
	"fmt"
	"log"
	"net/http"

	i "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/order_types"
	u "github.com/bhanupbalusu/custpreorderms/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type OrderTypesController struct {
	otsi s.OrderTypesServiceInterface
}

func NewOrderTypesController(otsi s.OrderTypesServiceInterface) i.OrderTypesControllerInterface {
	return &OrderTypesController{otsi}
}

func (otc *OrderTypesController) GetAll(ctx *fiber.Ctx) error {
	pd, err := otc.otsi.Get()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (otc *OrderTypesController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	fmt.Println(id)
	pd, err := otc.otsi.GetByID(id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (otc *OrderTypesController) Create(ctx *fiber.Ctx) error {
	var req m.OrderTypesModel
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	fmt.Println("------- Inside Handler Create Method Before Calling ProductService.Create -----------")
	fmt.Println(req)
	pid, err := otc.otsi.Create(&req)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pid)

}

func (otc *OrderTypesController) Update(ctx *fiber.Ctx) error {
	var req m.OrderTypesModel
	fmt.Println("-----------api/handler.Update Before calling c.BodyParser ----------")
	if err := ctx.BodyParser(&req); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	fmt.Println("-----------api/handler.Update Before calling h.ProductService.Update ----------")
	if err := otc.otsi.Update(&req); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Product failed to update",
			"error":   err.Error(),
		})
	}

	fmt.Println("-----------api/handler.Update Before calling final return----------")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
	})
}

func (otc *OrderTypesController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := otc.otsi.Delete(id); err != nil {
		log.Fatal(err)
	}
	return ctx.SendString("Product Is Deleted")
}
