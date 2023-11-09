package main

import (
	"log"
	"os"

	"github.com/tashima42/rock-paper-scissors-server/cmd"
)

func main() {
	app := cmd.NewRootCommand()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// game := rockpaperscissors.NewGame()

	// match := game.NewMatch(3)
	// player1 := rockpaperscissors.NewPlayer("player 1")
	// player2 := rockpaperscissors.NewPlayer("player 2")

	// if err := match.Join(player1); err != nil {
	// 	panic(err)
	// }
	// if err := match.Join(player2); err != nil {
	// 	panic(err)
	// }

	// if err := match.Start(); err != nil {
	// 	panic(err)
	// }

	// r, err := match.Play(rockpaperscissors.MoveTypePaper, player1.GetID())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r)
	// r, err = match.Play(rockpaperscissors.MoveTypeRock, player2.GetID())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r)
	// r, err = match.Play(rockpaperscissors.MoveTypeScissors, player1.GetID())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r)
	// r, err = match.Play(rockpaperscissors.MoveTypeRock, player2.GetID())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r)
}
