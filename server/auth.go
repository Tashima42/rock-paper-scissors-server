package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) authenticate(ctx *fiber.Ctx) error {
	auth := ctx.Cookies("auth")
	if auth == "" {
		return fiber.NewError(http.StatusUnauthorized, "player not registered")
	}
	claims, err := ParseJWT(c.JWTSecret, auth)
	if err != nil {
		return err
	}
	ctx.Locals("player", claims.Player)
	return ctx.Next()
}
