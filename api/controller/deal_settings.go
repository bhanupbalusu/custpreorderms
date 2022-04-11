package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	i "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/deal_settings"
	u "github.com/bhanupbalusu/custpreorderms/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type DealSettingsController struct {
	dssi s.DealSettingsServiceInterface
}

func NewDealSettingsController(pdsi s.DealSettingsServiceInterface) i.DealSettingsControllerInterface {
	return &DealSettingsController{pdsi}
}

func (dsc *DealSettingsController) GetAll(ctx *fiber.Ctx) error {
	pd, err := dsc.dssi.Get()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (dsc *DealSettingsController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	fmt.Println(id)
	pd, err := dsc.dssi.GetByID(id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pd)
}

func (dsc *DealSettingsController) Create(ctx *fiber.Ctx) error {
	var req m.DealSettingsModel
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	fmt.Println("------- Inside Handler Create Method Before Calling ProductService.Create -----------")
	fmt.Println(req)
	req.CreatedAt = time.Now()
	req.UpdatedAt = req.CreatedAt
	pid, err := dsc.dssi.Create(&req)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(pid)

}

func (dsc *DealSettingsController) Update(ctx *fiber.Ctx) error {
	var req m.DealSettingsModel
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
	req.UpdatedAt = time.Now()
	if err := dsc.dssi.Update(&req); err != nil {
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

func (dsc *DealSettingsController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := dsc.dssi.Delete(id); err != nil {
		log.Fatal(err)
	}
	return ctx.SendString("Product Is Deleted")
}
