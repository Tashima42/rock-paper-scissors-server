package server

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	rockpaperscissors "github.com/tashima42/rock-paper-scissors-server/rock-paper-scissors"
)

type Controller struct {
	JWTSecret []byte
	Game      *rockpaperscissors.Game
}

func (c *Controller) registerPlayer(ctx *fiber.Ctx) error {
	p := &rockpaperscissors.Player{}
	if err := json.Unmarshal(ctx.Body(), p); err != nil {
		return err
	}

	player := rockpaperscissors.NewPlayer(p.Name)

	jwt, err := NewJWT(c.JWTSecret, AuthClaims{Player: *player})
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour * 336),
		HTTPOnly: true,
		Secure:   true,
	})

	return ctx.JSON(map[string]interface{}{"player": player, "sessionToken": jwt})
}
