package server

import (
	"encoding/json"
	"net/http"

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

	player := ctx.Locals("player").(rockpaperscissors.Player)

	if err := m.Join(&player); err != nil {
		return err
	}

	return ctx.JSON(m)
}

func (c *Controller) joinMatch(ctx *fiber.Ctx) error {
	matchID := ctx.Params("id")
	m, err := c.Game.GetMatch(matchID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	player := ctx.Locals("player").(rockpaperscissors.Player)

	if err := m.Join(&player); err != nil {
		return err
	}
	return ctx.JSON(map[string]bool{"success": true})
}

func (c *Controller) listMatches(ctx *fiber.Ctx) error {
	m := c.Game.GetMatches()
	return ctx.JSON(m)
}

func (c *Controller) startMatch(ctx *fiber.Ctx) error {
	matchID := ctx.Params("id")
	m, err := c.Game.GetMatch(matchID)
	if err != nil {
		return err
	}
	player := ctx.Locals("player").(rockpaperscissors.Player)

	if ok := m.IsOneOfPlayers(player.ID); !ok {
		return fiber.NewError(http.StatusBadRequest, "player is not in this match")
	}

	if err := m.Start(); err != nil {
		return err
	}
	return ctx.JSON(map[string]bool{"success": true})
}

// func (c *Controller) playMatch(ctx *fiber.Ctx) error {
// }
