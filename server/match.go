package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	rockpaperscissors "github.com/tashima42/rock-paper-scissors-server/rock-paper-scissors"
)

type createMatchRequest struct {
	MaxScore int `json:"maxScore"`
}

func (c *Controller) createMatch(ctx *fiber.Ctx) error {
	cm := &createMatchRequest{}
	if err := json.Unmarshal(ctx.Body(), cm); err != nil {
		return err
	}
	m := c.Game.NewMatch(cm.MaxScore)

	player := ctx.Locals("player").(*rockpaperscissors.Player)

	if err := m.Join(player); err != nil {
		return err
	}

	return ctx.JSON(m)
}
