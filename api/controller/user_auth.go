package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	i "github.com/bhanupbalusu/custpreorderms/api/controller_interface"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/user_auth"
	u "github.com/bhanupbalusu/custpreorderms/pkg/util"
	e "github.com/bhanupbalusu/custpreorderms/pkg/util/user_auth"
	sec "github.com/bhanupbalusu/custpreorderms/security/user_auth"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/asaskevich/govalidator.v9"
)

type UserAuthController struct {
	uas s.UserAuthServiceInterface
}

func NewUserAuthController(uas s.UserAuthServiceInterface) i.UserAuthControllerInterface {
	return &UserAuthController{uas: uas}
}

func (uac *UserAuthController) SignUp(ctx *fiber.Ctx) error {
	fmt.Println(ctx)
	var newUser m.User
	err := ctx.BodyParser(&newUser)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(u.NewJError(err))
	}
	newUser.Email = u.NormalizeEmail(newUser.Email)
	if !govalidator.IsEmail(newUser.Email) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(u.NewJError(e.ErrInvalidEmail))
	}
	fmt.Println("---------Controller SignUp before calling Application.GetByEmail---------")
	exists, err := uac.uas.GetByEmail(newUser.Email)
	if err != nil {
		if strings.TrimSpace(newUser.Password) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(u.NewJError(e.ErrEmptyPassword))
		}
		newUser.Password, err = sec.EncryptPassword(newUser.Password)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(u.NewJError(err))
		}
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = newUser.CreatedAt
		newUser.Id = primitive.NewObjectID()
		fmt.Println("---------Controller SignUp before calling Application.Create---------")
		err = uac.uas.Create(&newUser)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(u.NewJError(err))
		}
		return ctx.
			Status(http.StatusCreated).
			JSON(newUser)
	}
	if exists != nil {
		err = e.ErrEmailAlreadyExists
	}
	return ctx.
		Status(http.StatusBadRequest).
		JSON(u.NewJError(err))
}

func (uac *UserAuthController) SignIn(ctx *fiber.Ctx) error {
	var input m.User
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(u.NewJError(err))
	}
	input.Email = u.NormalizeEmail(input.Email)
	if !govalidator.IsEmail(input.Email) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(u.NewJError(e.ErrInvalidEmail))
	}

	user, err := uac.uas.GetByEmail(input.Email)
	if err != nil {
		log.Printf("%s signin failed at GetByEmail: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(e.ErrInvalidCredentials))
	}
	fmt.Println("----- Controller Before calling VerifyPassword -----")
	err = sec.VerifyPassword(user.Password, input.Password)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(e.ErrInvalidCredentials))
	}
	fmt.Println("----- Controller Before calling GenerateNewToken -----")
	token, err := sec.GenerateNewToken(user.Id.Hex())
	if err != nil {
		log.Printf("%s signin failed at GenerateNewToken : %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"user":  user,
			"token": fmt.Sprintf("Bearer %s", token),
		})
}

func (uac *UserAuthController) GetUser(ctx *fiber.Ctx) error {
	fmt.Println("!!!!!! Controller - GetUser - AuthRequestWithId")
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(err))
	}
	fmt.Println("!!!!!! Controller - GetUser - GetByID - Before")
	user, err := uac.uas.GetByID(payload.Id)
	fmt.Println("!!!!!! Controller - GetUser - GetByID - After")
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

func (uac *UserAuthController) GetUsers(ctx *fiber.Ctx) error {
	users, err := uac.uas.Get()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(users)
}

func (uac *UserAuthController) PutUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(err))
	}
	var update m.User
	err = ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(u.NewJError(err))
	}
	update.Email = u.NormalizeEmail(update.Email)
	if !govalidator.IsEmail(update.Email) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(u.NewJError(e.ErrInvalidEmail))
	}
	fmt.Println(update.Email)
	fmt.Println("!!!!!! Controller - PutUser - GetByEmail - Before")
	exists, err := uac.uas.GetByEmail(update.Email)
	fmt.Println("!!!!!! Controller - PutUser - GetByEmail - After")
	fmt.Println(exists)
	if err != nil || exists.Id.Hex() == payload.Id {
		fmt.Println("------ Controller - PutUser - GetByID - Before -----")
		user, err := uac.uas.GetByID(payload.Id)
		fmt.Println("------ Controller - PutUser - GetByID - After -----")
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(u.NewJError(err))
		}
		user.Email = update.Email
		user.UpdatedAt = time.Now()
		fmt.Println(user)
		fmt.Println("------ Controller - PutUser - Update - Before -----")
		err = uac.uas.Update(user)
		fmt.Println("------ Controller - PutUser - Update - After -----")
		if err != nil {
			return ctx.
				Status(http.StatusUnprocessableEntity).
				JSON(u.NewJError(err))
		}
		return ctx.
			Status(http.StatusOK).
			JSON(user)

	}
	if exists != nil {
		fmt.Println("------- email already exists ------")
		err = e.ErrEmailAlreadyExists
	}
	fmt.Println("-------!!!!!!!!!!!!! Error befor return PutUser")
	return ctx.
		Status(http.StatusBadRequest).
		JSON(u.NewJError(err))
}

func (uac *UserAuthController) DeleteUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(u.NewJError(err))
	}
	err = uac.uas.Delete(payload.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(u.NewJError(err))
	}
	ctx.Set("Entity", payload.Id)
	return ctx.SendStatus(http.StatusNoContent)
}

func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	id := ctx.Params("id")
	if !primitive.IsValidObjectID(id) {
		return nil, e.ErrUnauthorized
	}
	fmt.Println("------!!!!!!!!!!!!! AuthRequestWithId-- Before ----------")
	token := ctx.Locals("user").(*jwt.Token)
	fmt.Println("------!!!!!!!!!!!!! AuthRequestWithId -- After  ------------")
	payload, err := sec.UserAuthParseToken(token.Raw)
	if err != nil {
		return nil, err
	}
	fmt.Println("------!!!!!!!!!!!!! AuthRequestWithId------------")
	if payload.Id != id || payload.Issuer != id {
		return nil, e.ErrUnauthorized
	}
	return payload, nil
}
