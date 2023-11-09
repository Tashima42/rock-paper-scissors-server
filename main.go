package main

import (
	"fmt"

	rockpaperscissors "github.com/tashima42/rock-paper-scissors-server/pkg/rock-paper-scissors"
)

func main() {
	game := rockpaperscissors.NewGame()

	match := game.NewMatch(3)
	player1 := rockpaperscissors.NewPlayer("player 1")
	player2 := rockpaperscissors.NewPlayer("player 2")

	if err := match.Join(player1); err != nil {
		panic(err)
	}
	if err := match.Join(player2); err != nil {
		panic(err)
	}

	if err := match.Start(); err != nil {
		panic(err)
	}

	r, err := match.Play(rockpaperscissors.MoveTypePaper, player1.GetID())
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	r, err = match.Play(rockpaperscissors.MoveTypeRock, player2.GetID())
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	r, err = match.Play(rockpaperscissors.MoveTypeScissors, player1.GetID())
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	r, err = match.Play(rockpaperscissors.MoveTypeRock, player2.GetID())
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
