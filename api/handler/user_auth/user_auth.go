package user_auth

import (
	"fmt"

	c "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	h "github.com/bhanupbalusu/custpreorderms/api/handler_interface"

	"github.com/gofiber/fiber/v2"
)

type userAuthRoutesHandler struct {
	uaci c.UserAuthControllerInterface
}

func NewAuthRoutesHandler(uaci c.UserAuthControllerInterface) h.UserAuthRoutesHandlerInterface {
	return &userAuthRoutesHandler{uaci}
}

func (r *userAuthRoutesHandler) SignUp(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	fmt.Println("---------Handler SignUp before calling Controller.SignUp---------")
	return r.uaci.SignUp(ctx)
}

func (r *userAuthRoutesHandler) SignIn(ctx *fiber.Ctx) error {
	return r.uaci.SignIn(ctx)
}

func (r *userAuthRoutesHandler) GetUser(ctx *fiber.Ctx) error {
	return r.uaci.GetUser(ctx)
}

func (r *userAuthRoutesHandler) GetUsers(ctx *fiber.Ctx) error {
	return r.uaci.GetUsers(ctx)
}

func (r *userAuthRoutesHandler) PutUser(ctx *fiber.Ctx) error {
	return r.uaci.PutUser(ctx)
}

func (r *userAuthRoutesHandler) DeleteUser(ctx *fiber.Ctx) error {
	return r.uaci.DeleteUser(ctx)
}

func (r *userAuthRoutesHandler) Install(app *fiber.App) {
	app.Post("/signup", r.SignUp)
	app.Post("/signin", r.SignIn)
	app.Get("/users", AuthRequired, r.GetUsers)
	app.Get("/users/:id", AuthRequired, r.GetUser)
	app.Put("/users/:id", AuthRequired, r.PutUser)
	app.Delete("/users/:id", AuthRequired, r.DeleteUser)
}
