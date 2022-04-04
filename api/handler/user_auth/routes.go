package user_auth

import (
	"net/http"

	u "github.com/bhanupbalusu/custpreorderms/pkg/util/user_auth"
	sec "github.com/bhanupbalusu/custpreorderms/security/user_auth"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

type Routes interface {
	Install(app *fiber.App)
}

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:    sec.JwtUserAuthSecretKey,
		SigningMethod: sec.JwtUserAuthSigningMethod,
		TokenLookup:   "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(http.StatusUnauthorized).
				JSON(u.NewJError(err))
		},
	})(ctx)
}
