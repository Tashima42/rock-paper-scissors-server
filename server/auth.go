package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) authenticate(ctx *fiber.Ctx) error {
	auth := ctx.Cookies("auth")
	if auth == "" {
		if authHeader, ok := ctx.GetReqHeaders()["Authorization"]; ok {
			auth = authHeader[0]
		}
		if auth == "" {
			return fiber.NewError(http.StatusUnauthorized, "missing session cookie or authorization header")
		}
	}
	claims, err := ParseJWT(c.JWTSecret, auth)
	if err != nil {
		return err
	}
	player := claims.Player
	ctx.Locals("player", player)
	return ctx.Next()
}
