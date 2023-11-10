package server

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	rockpaperscissors "github.com/tashima42/rock-paper-scissors-server/rock-paper-scissors"
)

func Serve(port string, jwtSecret []byte) error {
	c := Controller{
		JWTSecret: jwtSecret,
		Game:      rockpaperscissors.NewGame(),
	}
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Post("/player", c.registerPlayer)

	app.Use(c.authenticate)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]bool{"success": true})
	})
	app.Post("/match", c.createMatch)
	app.Post("/match/:id/join", c.joinMatch)
	app.Get("/match/all", c.listMatches)
	app.Post("/match/:id/start", c.startMatch)

	app.Get("/ws/:id", websocket.New(c.gameLoop, websocket.Config{Subprotocols: []string{"rockpaperscissors"}}))

	return app.Listen(":" + port)
}
