package server

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	rockpaperscissors "github.com/tashima42/rock-paper-scissors-server/rock-paper-scissors"
)

type PlayRequest struct {
	Move rockpaperscissors.MoveType
}

func (c *Controller) gameLoop(conn *websocket.Conn) {
	log.Println("connecting ws")
	player := conn.Locals("player").(rockpaperscissors.Player)
	log.Printf("player: %+v", player)
	matchID := conn.Params("id")
	log.Println("match id: " + matchID)
	m, err := c.Game.GetMatch(matchID)
	log.Printf("match: %+v", m)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	var (
		mt  int
		msg []byte
	)
	for {
		log.Println("for")
		if mt, msg, err = conn.ReadMessage(); err == nil {
			log.Println("new message")
			r := &PlayRequest{}
			if err := json.Unmarshal(msg, r); err != nil {
				errorMessage(mt, err, conn)
				break
			}
			result, err := m.Play(r.Move, player.ID)
			if err != nil {
				errorMessage(mt, err, conn)
				break
			}
			res, err := json.Marshal(map[string]int{"result": int(result)})
			if err != nil {
				errorMessage(mt, err, conn)
				break
			}
			if err = conn.WriteMessage(mt, res); err != nil {
				errorMessage(mt, err, conn)
				break
			}
			break
		}
	}
}

func errorMessage(messageType int, err error, conn *websocket.Conn) {
	conn.WriteMessage(messageType, []byte(err.Error()))
}
